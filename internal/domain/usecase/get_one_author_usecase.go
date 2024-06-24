package usecase

import (
	"context"
	"encoding/json"
	"github.com/andreis3/foodtosave-case/internal/domain/aggregate"
	"github.com/andreis3/foodtosave-case/internal/domain/cache"
	"github.com/andreis3/foodtosave-case/internal/domain/entity"
	"github.com/andreis3/foodtosave-case/internal/domain/observability"
	"github.com/andreis3/foodtosave-case/internal/domain/repository"
	"github.com/andreis3/foodtosave-case/internal/domain/uow"
	"github.com/andreis3/foodtosave-case/internal/util"
	"net/http"
	"time"
)

type GetOneAuthorUsecase struct {
	uow     uow.IUnitOfWork
	cache   cache.ICache
	metrics observability.IMetricAdapter
}

func NewGetOneAuthorAllBooksUsecase(uow uow.IUnitOfWork, cache cache.ICache, metrics observability.IMetricAdapter) *GetOneAuthorUsecase {
	return &GetOneAuthorUsecase{
		uow:     uow,
		cache:   cache,
		metrics: metrics,
	}
}
func (g *GetOneAuthorUsecase) GetOneAuthorAllBooks(id string) (aggregate.AuthorBookAggregate, *util.ValidationError) {
	start := time.Now()
	var agg = aggregate.AuthorBookAggregate{}
	result := g.cache.Get(id)
	if result != "" {
		_ = json.Unmarshal([]byte(result), &agg)

		end := time.Now()
		duration := float64(end.Sub(start).Milliseconds())
		g.metrics.HistogramOperationDuration(context.Background(), "getOne", "author", duration)
		return agg, nil
	}
	err := g.uow.Do(func(uow uow.IUnitOfWork) *util.ValidationError {
		authorRepository := uow.GetRepository(util.AUTH_REPOSITORY_KEY).(repository.IAuthorRepository)
		bookRepository := uow.GetRepository(util.BOOK_REPOSITORY_KEY).(repository.IBookRepository)
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
				Origin:      "GetOneAuthorUsecase.GetOneAuthorAllBooks",
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
		agg.AddAuthorAndBooks(authorResult.ToEntity(), booksEntity)
		return nil
	})
	if err != nil {
		return aggregate.AuthorBookAggregate{}, err
	}
	ttlCache := 10
	go g.cache.Set(id, agg, ttlCache)
	end := time.Now()
	duration := float64(end.Sub(start).Milliseconds())
	g.metrics.HistogramOperationDuration(context.Background(), "getOne", "author", duration)
	return agg, nil
}
