package command

import (
	"github.com/andreis3/foodtosave-case/internal/app/command"
	"github.com/andreis3/foodtosave-case/internal/domain/services"
	"github.com/andreis3/foodtosave-case/internal/infra/adapters/db"
	"github.com/andreis3/foodtosave-case/internal/infra/commons/logger"
	"github.com/andreis3/foodtosave-case/internal/infra/commons/observability"
	"github.com/andreis3/foodtosave-case/internal/infra/commons/uuid"
	"github.com/andreis3/foodtosave-case/internal/infra/repository/redis/cache"
	"github.com/andreis3/foodtosave-case/internal/infra/uow"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func MakeCreateAuthorWithBooksCommand(postgresDB db.IDatabase, redisDB db.IDatabase, prometheus observability.IMetricAdapter) command.ICreateAuthorCommand {
	log := logger.NewLogger()
	postgresPool := postgresDB.InstanceDB().(*pgxpool.Pool)
	redisClient := redisDB.InstanceDB().(*redis.Client)
	cache := cache.NewCache(redisClient, prometheus, log)
	uuidGenerator := uuid.NewUUID()
	unitOfWork := uow.NewProxyUnitOfWork(postgresPool, prometheus)
	createAuthorService := services.NewAuthorService(unitOfWork, cache, prometheus)
	createAuthorCommand := command.NewCreateAuthorCommand(createAuthorService, uuidGenerator)
	return createAuthorCommand
}
