package make_handler

import (
	"github.com/andreis3/foodtosave-case/internal/domain/observability"
	"github.com/andreis3/foodtosave-case/internal/infra/adapters/db"
	"github.com/andreis3/foodtosave-case/internal/infra/common/logger"

	"github.com/andreis3/foodtosave-case/internal/infra/common/uuid"
	"github.com/andreis3/foodtosave-case/internal/presentation/http/hanlders/authorhandler"
)

func FactoryCreateAuthorHandler(postgresDB db.IDatabase, redisDB db.IDatabase, prometheus observability.IMetricAdapter) authorhandler.CreateAuthorWithBooksHandler {
	log := logger.NewLogger()
	requestID := uuid.NewUUID()
	createGroupHandler := authorhandler.NewCreateAuthorWithBooksHandler(postgresDB, redisDB, prometheus, log, requestID)
	return *createGroupHandler
}
