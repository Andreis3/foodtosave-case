package main

import (
	"fmt"
	"github.com/andreis3/foodtosave-case/internal/infra/common/configs"
	"github.com/andreis3/foodtosave-case/internal/infra/common/logger"
	"github.com/andreis3/foodtosave-case/internal/infra/server"
	"github.com/andreis3/foodtosave-case/internal/util"
	"os"
)

func main() {
	log := logger.NewLogger()
	conf, err := configs.LoadConfig(".")
	if err != nil {
		log.ErrorText(fmt.Sprintf("NotificationErrors loading config: %s", err.Error()))
		os.Exit(util.EXIT_FAILURE)
	}
	server.Start(conf, log)
}
