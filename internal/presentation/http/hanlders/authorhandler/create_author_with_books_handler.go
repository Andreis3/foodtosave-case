package authorhandler

import (
	"context"
	"github.com/andreis3/foodtosave-case/internal/domain/observability"
	"github.com/andreis3/foodtosave-case/internal/domain/uuid"
	"github.com/andreis3/foodtosave-case/internal/infra/adapters/db"
	"github.com/andreis3/foodtosave-case/internal/infra/common/logger"

	"github.com/andreis3/foodtosave-case/internal/infra/dto"
	"github.com/andreis3/foodtosave-case/internal/infra/factory/command"
	"github.com/andreis3/foodtosave-case/internal/presentation/http/helpers"
	"net/http"
	"strings"
	"time"
)

type CreateAuthorWithBooksHandler struct {
	logger     logger.ILogger
	id         uuid.IUUID
	prometheus observability.IMetricAdapter
	postgres   db.IDatabase
	redis      db.IDatabase
}

func NewCreateAuthorWithBooksHandler(
	postgres db.IDatabase,
	redis db.IDatabase,
	prometheus observability.IMetricAdapter,
	logger logger.ILogger,
	id uuid.IUUID) *CreateAuthorWithBooksHandler {
	return &CreateAuthorWithBooksHandler{
		logger:     logger,
		id:         id,
		prometheus: prometheus,
		postgres:   postgres,
		redis:      redis,
	}
}

func (cgc *CreateAuthorWithBooksHandler) CreateAuthorWithBooks(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	requestID := cgc.id.Generate()
	createAuthorCommand := command.FactoryCreateAuthorWithBooksCommand(cgc.postgres, cgc.redis, cgc.prometheus)
	groupInputDTO, err := helpers.DecoderBodyRequest[*dto.AuthorInput](r)
	if err != nil {
		cgc.logger.ErrorJson("Create Author NotificationErrors",
			"REQUEST_ID", requestID,
			"CODE_ERROR", err.Code,
			"ORIGIN", err.Origin,
			"ERROR_MESSAGE", strings.Join(err.LogError, ", "))
		cgc.prometheus.CounterRequestHttpStatusCode(context.Background(), helpers.CREATE_AUTHOR_V1, err.Status)
		end := time.Now()
		duration := end.Sub(start).Milliseconds()
		cgc.prometheus.HistogramRequestDuration(context.Background(), helpers.CREATE_AUTHOR_V1, err.Status, float64(duration))
		helpers.ResponseError[[]string](w, err.Status, requestID, err.Code, err.ClientError)
		return
	}
	author, errCM := createAuthorCommand.Execute(*groupInputDTO)
	if errCM != nil {
		cgc.logger.ErrorJson("Create Group NotificationErrors",
			"REQUEST_ID", requestID,
			"CODE_ERROR", errCM.Code,
			"ORIGIN", errCM.Origin,
			"ERROR_MESSAGE", strings.Join(errCM.LogError, ", "))
		cgc.prometheus.CounterRequestHttpStatusCode(context.Background(), helpers.CREATE_AUTHOR_V1, errCM.Status)
		end := time.Now()
		duration := end.Sub(start).Milliseconds()
		cgc.prometheus.HistogramRequestDuration(context.Background(), helpers.CREATE_AUTHOR_V1, errCM.Status, float64(duration))
		helpers.ResponseError[[]string](w, errCM.Status, requestID, errCM.Code, errCM.ClientError)
		return
	}
	cgc.prometheus.CounterRequestHttpStatusCode(context.Background(), helpers.CREATE_AUTHOR_V1, http.StatusCreated)
	end := time.Now()
	duration := end.Sub(start).Milliseconds()
	cgc.prometheus.HistogramRequestDuration(context.Background(), helpers.CREATE_AUTHOR_V1, http.StatusCreated, float64(duration))
	helpers.ResponseSuccess[dto.AuthorOutput](w, requestID, http.StatusCreated, author)
}
