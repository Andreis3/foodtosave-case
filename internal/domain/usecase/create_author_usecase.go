package usecase

import (
	"context"
	"net/http"
	"time"

	"github.com/andreis3/foodtosave-case/internal/domain/aggregate"
	"github.com/andreis3/foodtosave-case/internal/domain/cache"
	"github.com/andreis3/foodtosave-case/internal/domain/observability"
	"github.com/andreis3/foodtosave-case/internal/domain/repository"
	"github.com/andreis3/foodtosave-case/internal/domain/uow"
	"github.com/andreis3/foodtosave-case/internal/util"
)

type CreateAuthorUseCase struct {
	uow     uow.IUnitOfWork
	cache   cache.ICache
	metrics observability.IMetricAdapter
}

func NewCreateAuthorWithBookUseCase(uow uow.IUnitOfWork, cache cache.ICache, metrics observability.IMetricAdapter) *CreateAuthorUseCase {
	return &CreateAuthorUseCase{
		uow:     uow,
		cache:   cache,
		metrics: metrics,
	}
}
func (c *CreateAuthorUseCase) CreateAuthorWithBooks(data aggregate.AuthorBookAggregate) (aggregate.AuthorBookAggregate, *util.ValidationError) {
	start := time.Now()
	aggValidate := data.Validate()
	if aggValidate.HasErrors() {
		return aggregate.AuthorBookAggregate{}, &util.ValidationError{
			Code:        "VBR-0001",
			Origin:      "CreateAuthorUseCase.CreateAuthorWithBooks",
			LogError:    aggValidate.ReturnErrors(),
			ClientError: aggValidate.ReturnErrors(),
			Status:      http.StatusBadRequest,
		}
	}
	err := c.uow.Do(func(uow uow.IUnitOfWork) *util.ValidationError {
		authorRepository := c.uow.GetRepository(util.AUTH_REPOSITORY_KEY).(repository.IAuthorRepository)
		bookRepository := c.uow.GetRepository(util.BOOK_REPOSITORY_KEY).(repository.IBookRepository)
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
	go c.cache.Set(data.Author.ID, data, ttlCache)
	end := time.Now()
	duration := float64(end.Sub(start).Milliseconds())
	c.metrics.HistogramOperationDuration(context.Background(), "create", "authors", duration)
	return data, nil
}
