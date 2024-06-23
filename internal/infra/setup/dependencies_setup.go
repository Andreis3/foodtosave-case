package setup

import (
	"github.com/andreis3/foodtosave-case/internal/infra/adapters/db"
	"github.com/andreis3/foodtosave-case/internal/infra/common/logger"
	"github.com/andreis3/foodtosave-case/internal/infra/common/observability"
	"github.com/andreis3/foodtosave-case/internal/infra/factory/handler"
	"github.com/andreis3/foodtosave-case/internal/infra/routes"
	"github.com/andreis3/foodtosave-case/internal/presentation/http/hanlders/authorhandler/authorroutes"
	"github.com/go-chi/chi/v5"
)

func SetupRoutesAndDependencies(mux *chi.Mux, postgres db.IDatabase, redis db.IDatabase, logger logger.ILogger) {
	registerRouter := routes.NewRegisterRoutes(logger)
	prometheus := observability.NewPrometheusAdapter()
	createAuthorHandler := make_handler.FactoryCreateAuthorHandler(postgres, redis, prometheus)
	getOneAuthorHandler := make_handler.FactoryGetOneAuthorAllBooksHandler(postgres, redis, prometheus)
	auhtorRouter := authorroutes.NewAuthorRoutes(createAuthorHandler, getOneAuthorHandler)
	routes.NewRoutes(mux, *registerRouter, *auhtorRouter).RegisterRoutes()
}
