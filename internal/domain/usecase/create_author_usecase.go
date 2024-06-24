package usecase

import (
	"context"
	"github.com/andreis3/foodtosave-case/internal/domain/aggregate"
	"github.com/andreis3/foodtosave-case/internal/domain/cache"
	"github.com/andreis3/foodtosave-case/internal/domain/observability"
	"github.com/andreis3/foodtosave-case/internal/domain/repository"
	"github.com/andreis3/foodtosave-case/internal/domain/uow"
	"github.com/andreis3/foodtosave-case/internal/util"
	"net/http"
	"time"
)

type CreateAuthorUsecase struct {
	uow     uow.IUnitOfWork
	cache   cache.ICache
	metrics observability.IMetricAdapter
}

func NewCreateAuthorWithBookUsecase(uow uow.IUnitOfWork, cache cache.ICache, metrics observability.IMetricAdapter) *CreateAuthorUsecase {
	return &CreateAuthorUsecase{
		uow:     uow,
		cache:   cache,
		metrics: metrics,
	}
}
func (cas *CreateAuthorUsecase) CreateAuthorWithBooks(data aggregate.AuthorBookAggregate) (aggregate.AuthorBookAggregate, *util.ValidationError) {
	start := time.Now()
	aggValidate := data.Validate()
	if aggValidate.HasErrors() {
		return aggregate.AuthorBookAggregate{}, &util.ValidationError{
			Code:        "VBR-0001",
			Origin:      "CreateAuthorUsecase.CreateAuthorWithBooks",
			LogError:    aggValidate.ReturnErrors(),
			ClientError: aggValidate.ReturnErrors(),
			Status:      http.StatusBadRequest,
		}
	}
	err := cas.uow.Do(func(uow uow.IUnitOfWork) *util.ValidationError {
		authorRepository := cas.uow.GetRepository(util.AUTH_REPOSITORY_KEY).(repository.IAuthorRepository)
		bookRepository := cas.uow.GetRepository(util.BOOK_REPOSITORY_KEY).(repository.IBookRepository)
		authorId, err := authorRepository.InsertAuthor(data.Author)
		if err != nil {
			return err
		}
		data.Author.ID = authorId
		for index := range data.Books {
			bookId, err := bookRepository.InsertBook(data.Books[index], data.Author.ID)
			if err != nil {
				return err
			}
			data.Books[index].ID = bookId
		}
		return nil
	})
	if err != nil {
		return aggregate.AuthorBookAggregate{}, err
	}
	ttlCache := 10
	go cas.cache.Set(data.Author.ID, data, ttlCache)
	end := time.Now()
	duration := float64(end.Sub(start).Milliseconds())
	cas.metrics.HistogramOperationDuration(context.Background(), "create", "authors", duration)
	return data, nil
}
