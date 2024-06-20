package uow

import (
	"github.com/andreis3/foodtosave-case/internal/infra/adapters/db"
	"github.com/andreis3/foodtosave-case/internal/infra/commons/observability"
	repo_group "github.com/andreis3/foodtosave-case/internal/infra/repository/postgres/author"
	"github.com/andreis3/foodtosave-case/internal/util"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewProxyUnitOfWork(pool *pgxpool.Pool, metric observability.IMetricAdapter) *UnitOfWork {
	uow := NewUnitOfWork(pool)
	uow.Register(util.AUTH_REPOSITORY_KEY, func(tx any) any {
		repo := repo_group.NewAuthorRepository(metric)
		repo.DB = db.New(tx.(pgx.Tx))
		return repo
	})
	return uow
}
