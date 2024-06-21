package main

import (
	"fmt"
	"github.com/andreis3/foodtosave-case/internal/infra/commons/configs"
	"github.com/andreis3/foodtosave-case/internal/infra/commons/logger"
	"github.com/andreis3/foodtosave-case/internal/infra/server"
	"github.com/andreis3/foodtosave-case/internal/util"
	"os"
)

func main() {
	log := logger.NewLogger()
	conf, err := configs.LoadConfig(".")
	if err != nil {
		log.ErrorText(fmt.Sprintf("Error loading config: %s", err.Error()))
		os.Exit(util.EXIT_FAILURE)
	}
	server.Start(conf, log)
}
