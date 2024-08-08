package authorhandler

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/andreis3/foodtosave-case/internal/domain/observability"
	"github.com/andreis3/foodtosave-case/internal/infra/adapters/db"
	"github.com/andreis3/foodtosave-case/internal/infra/common/logger"
	"github.com/andreis3/foodtosave-case/internal/infra/common/uuid"

	"github.com/andreis3/foodtosave-case/internal/infra/dto"
	"github.com/andreis3/foodtosave-case/internal/infra/factory/query"
	"github.com/andreis3/foodtosave-case/internal/presentation/http/helpers"
)

type GetOneAuthorAllBooksHandler struct {
	logger     logger.ILogger
	prometheus observability.IMetricAdapter
	id         uuid.IUUID
	postgres   db.IDatabase
	redis      db.IDatabase
}

func NewGetOneAuthorAllBooksHandler(
	postgres db.IDatabase,
	redis db.IDatabase,
	prometheus observability.IMetricAdapter,
	logger logger.ILogger,
	id uuid.IUUID) *GetOneAuthorAllBooksHandler {
	return &GetOneAuthorAllBooksHandler{
		logger:     logger,
		prometheus: prometheus,
		id:         id,
		postgres:   postgres,
		redis:      redis,
	}
}

func (h *GetOneAuthorAllBooksHandler) GetOneAuthorAllBooks(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	getOneAuthorQuery := query.FactoryGetOneAuthorAllBooksQuery(h.postgres, h.redis, h.prometheus)
	requestID := h.id.Generate()
	err := helpers.PathRouterValidate(r, helpers.ID)
	if err != nil {
		h.logger.ErrorJson("Get Author NotificationErrors",
			"REQUEST_ID", requestID,
			"CODE_ERROR", err.Code,
			"ORIGIN", err.Origin,
			"ERROR_MESSAGE", strings.Join(err.LogError, ", "))
		h.prometheus.CounterRequestHttpStatusCode(context.Background(), helpers.GET_AUTHOR_V1, err.Status)
		end := time.Now()
		duration := end.Sub(start).Milliseconds()
		h.prometheus.HistogramRequestDuration(context.Background(), helpers.GET_AUTHOR_V1, err.Status, float64(duration))
		helpers.ResponseError[[]string](w, err.Status, requestID, err.Code, err.ClientError)
		return
	}
	id := r.PathValue("id")
	group, err := getOneAuthorQuery.Execute(id)
	if err != nil {
		h.logger.ErrorJson("Select One Group NotificationErrors",
			"REQUEST_ID", requestID,
			"CODE_ERROR", err.Code,
			"ORIGIN", err.Origin,
			"ERROR_MESSAGE", strings.Join(err.LogError, ", "))
		h.prometheus.CounterRequestHttpStatusCode(context.Background(), helpers.GET_AUTHOR_V1, err.Status)
		end := time.Now()
		duration := end.Sub(start).Milliseconds()
		h.prometheus.HistogramRequestDuration(context.Background(), helpers.GET_AUTHOR_V1, err.Status, float64(duration))
		helpers.ResponseError[[]string](w, err.Status, requestID, err.Code, err.ClientError)
		return
	}
	h.prometheus.CounterRequestHttpStatusCode(context.Background(), helpers.GET_AUTHOR_V1, http.StatusOK)
	end := time.Now()
	duration := end.Sub(start).Milliseconds()
	h.prometheus.HistogramRequestDuration(context.Background(), helpers.GET_AUTHOR_V1, http.StatusOK, float64(duration))
	helpers.ResponseSuccess[dto.AuthorOutput](w, requestID, http.StatusOK, group)
}
