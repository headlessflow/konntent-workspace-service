package main

import (
	"context"
	"konntent-workspace-service/pkg/constants"
	"konntent-workspace-service/pkg/nrclient"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := initLogger()

	env := os.Getenv(constants.ConfigEnvKey)
	conf, cErr := initConfig(env)
	if cErr != nil {
		logger.Error(cErr)
		return
	}

	app, err := boot(logger, conf.Application)
	if err != nil {
		logger.Fatalf("Something went wrong while utilizing the server. %v", err)
	}
	sv := initServer(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	go log.Fatal(sv.Listen(":" + port))

	graceful(logger, sv, app.nrInstance)
}

func graceful(l *logrus.Logger, a *fiber.App, nr nrclient.NewRelicInstance) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	nr.Application().Shutdown(constants.AppGracefulTimeout * time.Second)
	_, cancel := context.WithTimeout(context.Background(), constants.AppGracefulTimeout*time.Second)
	defer cancel()
	if err := a.Shutdown(); err != nil {
		l.Fatal(err)
	}
}
