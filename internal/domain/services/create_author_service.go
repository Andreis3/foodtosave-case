package services

import (
	"context"
	"github.com/andreis3/foodtosave-case/internal/domain/aggregate"
	"github.com/andreis3/foodtosave-case/internal/infra/adapters/uuid"
	"github.com/andreis3/foodtosave-case/internal/infra/commons/observability"
	"github.com/andreis3/foodtosave-case/internal/infra/repository/postgres/author"
	"github.com/andreis3/foodtosave-case/internal/infra/repository/postgres/book"
	"github.com/andreis3/foodtosave-case/internal/infra/repository/redis/cache"
	"github.com/andreis3/foodtosave-case/internal/infra/uow"
	"github.com/andreis3/foodtosave-case/internal/interfaces/http/hanlders/author/dto"
	"github.com/andreis3/foodtosave-case/internal/util"
	"net/http"
	"time"
)

type CreateAuthorService struct {
	uow     uow.IUnitOfWork
	id      uuid.IUUID
	cache   cache.ICache
	metrics observability.IMetricAdapter
	output  dto.AuthorOutput
}

func NewAuthorService(uow uow.IUnitOfWork, id uuid.IUUID, cache cache.ICache, metrics observability.IMetricAdapter) *CreateAuthorService {
	return &CreateAuthorService{
		uow:     uow,
		id:      id,
		cache:   cache,
		metrics: metrics,
	}
}
func (cas *CreateAuthorService) CreateAuthor(data aggregate.AuthorBookAggregate) (dto.AuthorOutput, *util.ValidationError) {
	start := time.Now()
	authorEntity := data.Author
	booksEntity := data.Books
	authorEntity.ID = cas.id.Generate()
	authorValidate := authorEntity.Validate()
	for index := range booksEntity {
		booksEntity[index].ID = cas.id.Generate()
		authorValidate.MergeErrors(index, "books", booksEntity[index].Validate())
	}

	if authorValidate.HasErrors() {
		return dto.AuthorOutput{}, &util.ValidationError{
			Code:        "VBR-0001",
			Origin:      "CreateAuthorService.CreateAuthor",
			LogError:    authorValidate.ReturnErrors(),
			ClientError: authorValidate.ReturnErrors(),
			Status:      http.StatusBadRequest,
		}
	}
	err := cas.uow.Do(func(uow uow.IUnitOfWork) *util.ValidationError {
		authorRepository := cas.uow.GetRepository(util.AUTH_REPOSITORY_KEY).(author.IAuthorRepository)
		bookRepository := cas.uow.GetRepository(util.BOOK_REPOSITORY_KEY).(book.IBookRepository)
		_, err := authorRepository.InsertAuthor(authorEntity)
		if err != nil {
			return err
		}
		for _, book := range booksEntity {
			_, err = bookRepository.InsertBook(book, authorEntity.ID)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return dto.AuthorOutput{}, err
	}
	output := cas.output.MapperToAggregateAuthor(data.SetIDS(authorEntity, booksEntity))
	ttlCache := 10
	go cas.cache.SetNX(output.ID, output, ttlCache)
	end := time.Now()
	duration := float64(end.Sub(start).Milliseconds())
	cas.metrics.HistogramOperationDuration(context.Background(), "create", "authors", duration)
	return output, nil
}
