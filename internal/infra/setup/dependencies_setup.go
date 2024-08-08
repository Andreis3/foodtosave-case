package setup

import (
	"github.com/go-chi/chi/v5"

	"github.com/andreis3/foodtosave-case/internal/infra/adapters/db"
	"github.com/andreis3/foodtosave-case/internal/infra/common/logger"
	"github.com/andreis3/foodtosave-case/internal/infra/common/observability"
	"github.com/andreis3/foodtosave-case/internal/infra/factory/handler"
	"github.com/andreis3/foodtosave-case/internal/infra/routes"
	"github.com/andreis3/foodtosave-case/internal/presentation/http/hanlders/authorhandler/authorroutes"
)

func RoutesAndDependencies(mux *chi.Mux, postgres db.IDatabase, redis db.IDatabase, logger logger.ILogger) {
	prometheus := observability.NewPrometheusAdapter()
	createAuthorHandler := make_handler.FactoryCreateAuthorHandler(postgres, redis, prometheus)
	getOneAuthorHandler := make_handler.FactoryGetOneAuthorAllBooksHandler(postgres, redis, prometheus)
	authorRoutes := authorroutes.NewAuthorRoutes(createAuthorHandler, getOneAuthorHandler)
	routes.NewRegisterRoutes(mux, *authorRoutes, logger).RegisterRoutes()
}
