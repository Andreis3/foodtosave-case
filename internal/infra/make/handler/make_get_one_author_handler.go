package make_handler

import (
	"github.com/andreis3/foodtosave-case/internal/infra/adapters/db"
	"github.com/andreis3/foodtosave-case/internal/infra/adapters/uuid"
	"github.com/andreis3/foodtosave-case/internal/infra/commons/logger"
	"github.com/andreis3/foodtosave-case/internal/infra/commons/observability"
	"github.com/andreis3/foodtosave-case/internal/interfaces/http/hanlders/author"
)

func MakeGetOneGroupHandler(postgresDB db.IDatabase, redisDB db.IDatabase, prometheus observability.IMetricAdapter) handler_author.IGetOneAuthorHandler {
	log := logger.NewLogger()
	requestID := uuid.NewUUID()
	searchGroupController := handler_author.NewGetOneGroupHandler(postgresDB, redisDB, prometheus, log, requestID)
	return searchGroupController
}
