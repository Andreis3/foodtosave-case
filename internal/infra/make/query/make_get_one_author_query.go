package query

import (
	"github.com/andreis3/foodtosave-case/internal/app/query"
	"github.com/andreis3/foodtosave-case/internal/domain/services"
	"github.com/andreis3/foodtosave-case/internal/infra/adapters/db"
	"github.com/andreis3/foodtosave-case/internal/infra/commons/logger"
	"github.com/andreis3/foodtosave-case/internal/infra/commons/observability"
	"github.com/andreis3/foodtosave-case/internal/infra/repository/redis/cache"
	"github.com/andreis3/foodtosave-case/internal/infra/uow"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func MakeGetOneAuthorQuery(postgresDB db.IDatabase, redisDB db.IDatabase, prometheus observability.IMetricAdapter) query.IGetOneAuthorQuery {
	log := logger.NewLogger()
	postgresPool := postgresDB.InstanceDB().(*pgxpool.Pool)
	redisClient := redisDB.InstanceDB().(*redis.Client)
	cache := cache.NewCache(redisClient, prometheus, log)
	unitOfWork := uow.NewProxyUnitOfWork(postgresPool, prometheus)
	getOneAuthorService := services.NewGetOneService(unitOfWork, cache, prometheus)
	getOneAuthorCommand := query.NewGetAuthorQuery(getOneAuthorService)
	return getOneAuthorCommand
}
