package proxy

import (
	"github.com/andreis3/foodtosave-case/internal/infra/adapters/db"
	"github.com/andreis3/foodtosave-case/internal/infra/commons/logger"
	"github.com/andreis3/foodtosave-case/internal/infra/commons/observability"
	"github.com/andreis3/foodtosave-case/internal/infra/make/handler"
	"github.com/andreis3/foodtosave-case/internal/infra/routes"
	"github.com/andreis3/foodtosave-case/internal/interfaces/http/hanlders/authorhandler/routes"
	"github.com/go-chi/chi/v5"
)

func ProxyDependency(mux *chi.Mux, postgres db.IDatabase, redis db.IDatabase, logger logger.ILogger) {
	registerRouter := routes.NewRegisterRoutes(logger)
	prometheus := observability.NewPrometheusAdapter()
	createAuthorHandler := make_handler.MakeCreateAuthorHandler(postgres, redis, prometheus)
	getOneAuthorHandler := make_handler.MakeGetOneAuthorAllBooksHandler(postgres, redis, prometheus)
	auhtorRouter := author_routes.NewAuthorRoutes(createAuthorHandler, getOneAuthorHandler)
	routes.NewRoutes(mux, *registerRouter, *auhtorRouter).RegisterRoutes()
}
