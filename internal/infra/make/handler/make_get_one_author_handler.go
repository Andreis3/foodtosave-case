package make_handler

import (
	"github.com/andreis3/foodtosave-case/internal/infra/adapters/db"
	"github.com/andreis3/foodtosave-case/internal/infra/adapters/uuid"
	"github.com/andreis3/foodtosave-case/internal/infra/commons/logger"
	"github.com/andreis3/foodtosave-case/internal/infra/commons/observability"
	"github.com/andreis3/foodtosave-case/internal/interfaces/http/hanlders/authorhandler"
)

func MakeGetOneAuthorAllBooksHandler(postgresDB db.IDatabase, redisDB db.IDatabase, prometheus observability.IMetricAdapter) authorhandler.GetOneAuthorAllBooksHandler {
	log := logger.NewLogger()
	requestID := uuid.NewUUID()
	searchGroupController := authorhandler.NewGetOneAuthorAllBooksHandler(postgresDB, redisDB, prometheus, log, requestID)
	return *searchGroupController
}
