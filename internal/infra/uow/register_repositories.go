package uow

import (
	"github.com/andreis3/foodtosave-case/internal/domain/observability"
	"github.com/andreis3/foodtosave-case/internal/infra/adapters/db/postgres"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/andreis3/foodtosave-case/internal/infra/repository/postgres/author"
	"github.com/andreis3/foodtosave-case/internal/infra/repository/postgres/book"
	"github.com/andreis3/foodtosave-case/internal/util"
)

func NewProxyUnitOfWork(pool *pgxpool.Pool, metric observability.IMetricAdapter) *UnitOfWork {
	uow := NewUnitOfWork(pool)
	uow.Register(util.AUTH_REPOSITORY_KEY, func(tx any) any {
		repo := author.NewAuthorRepository(metric)
		repo.DB = postgres.New(tx.(pgx.Tx))
		return repo
	})
	uow.Register(util.BOOK_REPOSITORY_KEY, func(tx any) any {
		repo := book.NewBookRepository(metric)
		repo.DB = postgres.New(tx.(pgx.Tx))
		return repo
	})
	return uow
}
