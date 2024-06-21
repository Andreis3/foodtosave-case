package make_handler

import (
	"github.com/andreis3/foodtosave-case/internal/infra/adapters/db"
	"github.com/andreis3/foodtosave-case/internal/infra/adapters/uuid"
	"github.com/andreis3/foodtosave-case/internal/infra/commons/logger"
	"github.com/andreis3/foodtosave-case/internal/infra/commons/observability"
	"github.com/andreis3/foodtosave-case/internal/interfaces/http/hanlders/author"
)

func MakeCreateAuthorHandler(postgresDB db.IDatabase, redisDB db.IDatabase, prometheus observability.IMetricAdapter) handler_author.ICreateAuthorHandler {
	log := logger.NewLogger()
	requestID := uuid.NewUUID()
	createGroupHandler := handler_author.NewCreateAuthorHandler(postgresDB, redisDB, prometheus, log, requestID)
	return createGroupHandler
}
