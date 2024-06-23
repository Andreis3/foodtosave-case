package usecase

import (
	"context"
	"github.com/andreis3/foodtosave-case/internal/domain/aggregate"
	"github.com/andreis3/foodtosave-case/internal/infra/common/observability"
	"github.com/andreis3/foodtosave-case/internal/infra/dto"
	"github.com/andreis3/foodtosave-case/internal/infra/repository/postgres/author"
	"github.com/andreis3/foodtosave-case/internal/infra/repository/postgres/book"
	"github.com/andreis3/foodtosave-case/internal/infra/repository/redis/cache"
	"github.com/andreis3/foodtosave-case/internal/infra/uow"
	"github.com/andreis3/foodtosave-case/internal/util"
	"net/http"
	"time"
)

type CreateAuthorWithBookUsecase struct {
	uow     uow.IUnitOfWork
	cache   cache.ICache
	metrics observability.IMetricAdapter
}

func NewCreateAuthorWithBookUsecase(uow uow.IUnitOfWork, cache cache.ICache, metrics observability.IMetricAdapter) *CreateAuthorWithBookUsecase {
	return &CreateAuthorWithBookUsecase{
		uow:     uow,
		cache:   cache,
		metrics: metrics,
	}
}
func (cas *CreateAuthorWithBookUsecase) CreateAuthorWithBooks(data aggregate.AuthorBookAggregate) (dto.AuthorOutput, *util.ValidationError) {
	start := time.Now()
	aggValidate := data.Validate()
	if aggValidate.HasErrors() {
		return dto.AuthorOutput{}, &util.ValidationError{
			Code:        "VBR-0001",
			Origin:      "CreateAuthorWithBookUsecase.CreateAuthorWithBooks",
			LogError:    aggValidate.ReturnErrors(),
			ClientError: aggValidate.ReturnErrors(),
			Status:      http.StatusBadRequest,
		}
	}
	err := cas.uow.Do(func(uow uow.IUnitOfWork) *util.ValidationError {
		authorRepository := cas.uow.GetRepository(util.AUTH_REPOSITORY_KEY).(author.IAuthorRepository)
		bookRepository := cas.uow.GetRepository(util.BOOK_REPOSITORY_KEY).(book.IBookRepository)
		_, err := authorRepository.InsertAuthor(data.Author)
		if err != nil {
			return err
		}
		for _, book := range data.Books {
			_, err = bookRepository.InsertBook(book, data.Author.ID)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return dto.AuthorOutput{}, err
	}
	output := data.MapperToDtoOutput()
	ttlCache := 10
	go cas.cache.SetNX(output.ID, output, ttlCache)
	end := time.Now()
	duration := float64(end.Sub(start).Milliseconds())
	cas.metrics.HistogramOperationDuration(context.Background(), "create", "authors", duration)
	return output, nil
}
