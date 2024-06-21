package handler_author

import (
	"context"
	"github.com/andreis3/foodtosave-case/internal/infra/adapters/db"
	"github.com/andreis3/foodtosave-case/internal/infra/adapters/uuid"
	"github.com/andreis3/foodtosave-case/internal/infra/commons/logger"
	"github.com/andreis3/foodtosave-case/internal/infra/commons/observability"
	"github.com/andreis3/foodtosave-case/internal/infra/make/command"
	"github.com/andreis3/foodtosave-case/internal/interfaces/http/hanlders/author/dto"
	"github.com/andreis3/foodtosave-case/internal/interfaces/http/helpers"
	"net/http"
	"strings"
	"time"
)

type CreateAuthorHandler struct {
	logger     logger.ILogger
	id         uuid.IUUID
	prometheus observability.IMetricAdapter
	postgres   db.IDatabase
	redis      db.IDatabase
}

func NewCreateAuthorHandler(
	postgres db.IDatabase,
	redis db.IDatabase,
	prometheus observability.IMetricAdapter,
	logger logger.ILogger,
	id uuid.IUUID) *CreateAuthorHandler {
	return &CreateAuthorHandler{
		logger:     logger,
		id:         id,
		prometheus: prometheus,
		postgres:   postgres,
		redis:      redis,
	}
}

func (cgc *CreateAuthorHandler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	requestID := cgc.id.Generate()
	createAuthorCommand := command.MakeCreateAuthorCommand(cgc.postgres, cgc.redis, cgc.prometheus)
	groupInputDTO, err := helpers.DecoderBodyRequest[*dto.AuthorInput](r)
	if err != nil {
		cgc.logger.ErrorJson("Create Author Error",
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
		cgc.logger.ErrorJson("Create Group Error",
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
