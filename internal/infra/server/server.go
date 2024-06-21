package server

import (
	"errors"
	"fmt"
	"github.com/andreis3/foodtosave-case/internal/infra/adapters/db"
	"github.com/andreis3/foodtosave-case/internal/infra/commons/configs"
	"github.com/andreis3/foodtosave-case/internal/infra/commons/logger"
	"github.com/andreis3/foodtosave-case/internal/infra/proxy"
	"github.com/andreis3/foodtosave-case/internal/util"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
)

func Start(conf *configs.Conf, log *logger.Logger) {
	start := time.Now()
	mux := chi.NewRouter()
	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", conf.ServerPort),
		Handler: mux,
	}
	pool := db.NewPostgresDB(*conf)
	redis := db.NewRedis(*conf)
	go func() {
		proxy.ProxyDependency(mux, pool, redis, log)
		end := time.Now()
		ms := end.Sub(start).Milliseconds()
		log.InfoText(fmt.Sprintf("Server started in %d ms", ms))
		log.InfoText(fmt.Sprintf("Start server on port %s", conf.ServerPort))
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.ErrorText(fmt.Sprintf("Error starting server: %s", err.Error()))
			os.Exit(util.EXIT_FAILURE)
		}
	}()
	gracefulShutdown(server, pool, redis, log)
}
