package server

import (
	"context"
	"fmt"
	"github.com/andreis3/foodtosave-case/internal/infra/adapters/db"
	"github.com/andreis3/foodtosave-case/internal/infra/commons/logger"
	"github.com/andreis3/foodtosave-case/internal/util"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func gracefulShutdown(server *http.Server, pool *db.Postgres, redis *db.Redis, log *logger.Logger) {
	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-shutdownSignal
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.InfoText("Initiating graceful shutdown...")
	if err := server.Shutdown(ctx); err != nil {
		log.ErrorText(fmt.Sprintf("Error during server shutdown: %s", err.Error()))
	}
	log.InfoText("Closing postgres connection...")
	pool.Close()
	log.InfoText("Closing redis connection...")
	redis.Close()
	log.InfoText("Shutdown complete exit code 0...")
	os.Exit(util.EXIT_SUCCESS)
}
