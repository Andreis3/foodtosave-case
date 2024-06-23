package services

import (
	"context"
	"encoding/json"
	"github.com/andreis3/foodtosave-case/internal/domain/aggregate"
	"github.com/andreis3/foodtosave-case/internal/domain/entity"
	"github.com/andreis3/foodtosave-case/internal/infra/commons/observability"
	"github.com/andreis3/foodtosave-case/internal/infra/repository/postgres/author"
	"github.com/andreis3/foodtosave-case/internal/infra/repository/postgres/book"
	"github.com/andreis3/foodtosave-case/internal/infra/repository/redis/cache"
	"github.com/andreis3/foodtosave-case/internal/infra/uow"
	"github.com/andreis3/foodtosave-case/internal/interfaces/http/hanlders/authorhandler/dto"
	"github.com/andreis3/foodtosave-case/internal/util"
	"net/http"
	"time"
)

type GetOneAuthorAllBooksService struct {
	uow     uow.IUnitOfWork
	cache   cache.ICache
	metrics observability.IMetricAdapter
}

func NewGetOneService(uow uow.IUnitOfWork, cache cache.ICache, metrics observability.IMetricAdapter) *GetOneAuthorAllBooksService {
	return &GetOneAuthorAllBooksService{
		uow:     uow,
		cache:   cache,
		metrics: metrics,
	}
}
func (g *GetOneAuthorAllBooksService) GetOneAuthorAllBooks(id string) (dto.AuthorOutput, *util.ValidationError) {
	start := time.Now()
	aggregateAuthor := new(aggregate.AuthorBookAggregate)
	result, err := g.cache.Get(id)
	if err == nil && result != "" {
		var output dto.AuthorOutput
		_ = json.Unmarshal([]byte(result), &output)
		end := time.Now()
		duration := float64(end.Sub(start).Milliseconds())
		g.metrics.HistogramOperationDuration(context.Background(), "getOne", "author", duration)
		return output, nil
	}
	err = g.uow.Do(func(uow uow.IUnitOfWork) *util.ValidationError {
		authorRepository := uow.GetRepository(util.AUTH_REPOSITORY_KEY).(author.IAuthorRepository)
		bookRepository := uow.GetRepository(util.BOOK_REPOSITORY_KEY).(book.IBookRepository)
		authorResult, err := authorRepository.SelectOneAuthorByID(id)
		if err != nil {
			return err
		}
		if authorResult.ID == nil {
			end := time.Now()
			duration := float64(end.Sub(start).Milliseconds())
			g.metrics.HistogramOperationDuration(context.Background(), "getOne", "author", duration)
			return &util.ValidationError{
				Code:        "VBR-0002",
				Origin:      "GetOneAuthorAllBooksService.GetOneAuthorAllBooks",
				LogError:    []string{"Author not found"},
				ClientError: []string{"Author not found"},
				Status:      http.StatusNotFound,
			}
		}
		booksResult, err := bookRepository.SelectAllBooksByAuthorID(id)
		if err != nil {
			return err
		}
		var booksEntity []entity.Book
		for _, book := range booksResult {
			booksEntity = append(booksEntity, book.ToEntity())
		}
		aggregateAuthor.AddAuthorAndBooks(authorResult.ToEntity(), booksEntity)
		return nil
	})
	if err != nil {
		return dto.AuthorOutput{}, err
	}
	output := aggregateAuthor.MapperToDtoOutput()
	go g.cache.SetNX(id, output, 10)
	end := time.Now()
	duration := float64(end.Sub(start).Milliseconds())
	g.metrics.HistogramOperationDuration(context.Background(), "getOne", "author", duration)
	return output, nil
}
