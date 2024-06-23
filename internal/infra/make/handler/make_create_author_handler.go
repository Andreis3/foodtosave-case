package make_handler

import (
	"github.com/andreis3/foodtosave-case/internal/infra/adapters/db"
	"github.com/andreis3/foodtosave-case/internal/infra/adapters/uuid"
	"github.com/andreis3/foodtosave-case/internal/infra/commons/logger"
	"github.com/andreis3/foodtosave-case/internal/infra/commons/observability"
	"github.com/andreis3/foodtosave-case/internal/interfaces/http/hanlders/authorhandler"
)

func MakeCreateAuthorHandler(postgresDB db.IDatabase, redisDB db.IDatabase, prometheus observability.IMetricAdapter) authorhandler.CreateAuthorWithBooksHandler {
	log := logger.NewLogger()
	requestID := uuid.NewUUID()
	createGroupHandler := authorhandler.NewCreateAuthorWithBooksHandler(postgresDB, redisDB, prometheus, log, requestID)
	return *createGroupHandler
}
