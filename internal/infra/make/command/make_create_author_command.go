package command

import (
	"github.com/andreis3/foodtosave-case/internal/app/command"
	"github.com/andreis3/foodtosave-case/internal/domain/services"
	"github.com/andreis3/foodtosave-case/internal/infra/adapters/db"
	"github.com/andreis3/foodtosave-case/internal/infra/adapters/uuid"
	"github.com/andreis3/foodtosave-case/internal/infra/commons/observability"
	"github.com/andreis3/foodtosave-case/internal/infra/repository/redis/cache"
	"github.com/andreis3/foodtosave-case/internal/infra/uow"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func MakeCreateAuthorCommand(postgresDB db.IDatabase, redisDB db.IDatabase, prometheus observability.IMetricAdapter) command.ICreateAuthorCommand {
	postgresPool := postgresDB.InstanceDB().(*pgxpool.Pool)
	redisClient := redisDB.InstanceDB().(*redis.Client)
	cache := cache.NewCache(redisClient, prometheus)
	id := uuid.NewUUID()
	unitOfWork := uow.NewProxyUnitOfWork(postgresPool, prometheus)
	createAuthorService := services.NewAuthorService(unitOfWork, id, cache, prometheus)
	createAuthorCommand := command.NewCreateGroupCommand(createAuthorService)
	return createAuthorCommand
}
