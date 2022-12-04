package main

import (
	"konntent-workspace-service/pkg/dummyclient"
	"konntent-workspace-service/pkg/rabbit"

	"github.com/sirupsen/logrus"
)

type server struct {
	logger      *logrus.Logger
	dummyClient dummyclient.Client
	mqInstance  rabbit.ConsumerInstance
}

func initLogger() *logrus.Logger {
	return logrus.New()
}
