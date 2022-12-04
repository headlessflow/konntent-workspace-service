package main

import (
	"context"
	"konntent-workspace-service/pkg/constants"
	"konntent-workspace-service/pkg/rabbit"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	env := os.Getenv(constants.ConfigEnvKey)
	conf, err := initConfig(env)
	if err != nil {
		panic("<consumer> " + err.Error())
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})
	l := initLogger()

	l.Println("Starting " + constants.AppConsumerName + "...")

	app, err := boot(l, conf.Consumer)
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		defer wg.Done()
		app.mqInstance.Consume(context.Background())
	}()

	<-app.mqInstance.Handler().Status()
	wg.Wait()

	graceful(app.mqInstance)
}

func graceful(mq rabbit.ConsumerInstance) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	_, cancel := context.WithTimeout(context.Background(), constants.AppGracefulTimeout*time.Second)
	defer cancel()
	defer mq.Close()
}
