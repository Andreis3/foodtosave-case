package make_handler

import (
	"github.com/andreis3/foodtosave-case/internal/infra/adapters/db"
	"github.com/andreis3/foodtosave-case/internal/infra/common/logger"
	"github.com/andreis3/foodtosave-case/internal/infra/common/observability"
	"github.com/andreis3/foodtosave-case/internal/infra/common/uuid"
	"github.com/andreis3/foodtosave-case/internal/interfaces/http/hanlders/authorhandler"
)

func FactoryGetOneAuthorAllBooksHandler(postgresDB db.IDatabase, redisDB db.IDatabase, prometheus observability.IMetricAdapter) authorhandler.GetOneAuthorAllBooksHandler {
	log := logger.NewLogger()
	requestID := uuid.NewUUID()
	searchGroupController := authorhandler.NewGetOneAuthorAllBooksHandler(postgresDB, redisDB, prometheus, log, requestID)
	return *searchGroupController
}
