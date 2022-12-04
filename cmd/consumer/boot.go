package main

import (
	"fmt"
	"konntent-workspace-service/configs/app"
	"konntent-workspace-service/configs/consumer"
	"konntent-workspace-service/internal/listener"
	consumerservice "konntent-workspace-service/internal/listener/consumer"
	"konntent-workspace-service/pkg/constants"
	"konntent-workspace-service/pkg/dummyclient"
	"konntent-workspace-service/pkg/event"
	"konntent-workspace-service/pkg/eventmanager"
	"konntent-workspace-service/pkg/httpclient"
	"konntent-workspace-service/pkg/rabbit"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func boot(l *logrus.Logger, appConf consumer.Application) (*server, error) {
	var (
		httpClient         = httpclient.NewHTTPClient()
		dummyClient        = dummyclient.NewClient(initDummyClientConfig(appConf.Client), httpClient)
		mqPreProducer      = rabbit.NewPreProducer(l, app.RabbitQueueSettings(appConf.Rabbit.QueueSettings))
		mqConnectionClient = rabbit.NewClientConnector(l)
		mqConsumer         = rabbit.NewMessagingClient(l, mqConnectionClient, mqPreProducer)
	)

	var (
		consumerService     = consumerservice.NewDummyConsumerService(l, dummyClient)
		eventHandlerFactory = listener.NewEventHandlerFactory(consumerService)
		eventManager        = eventmanager.NewEventManager(eventHandlerFactory, event.NewEventCreator())
		customHandler       = listener.NewCustomHandler(l, eventManager)
		consumerManager     = rabbit.NewConsumerManager(l, customHandler)
		consumerGroup       = rabbit.NewSyncHandler(l, consumerManager, appConf.Rabbit.QueueSettings.MaxRetries)
	)

	// Open broker connection
	err := mqConsumer.ConnectToBroker(appConf.Rabbit.URL)
	if err != nil {
		return nil, err
	}

	return &server{
		logger:      l,
		dummyClient: dummyClient,
		mqInstance:  rabbit.NewConsumerInstance(l, mqConsumer, consumerGroup),
	}, nil
}

func initDummyClientConfig(dummyConf consumer.ClientConfig) dummyclient.Config {
	return dummyclient.Config{
		URL: dummyConf.URL,
	}
}

func initConfig[T string | constants.AppEnvironment](env T) (*consumer.Configs, error) {
	if env == "" {
		env = T(constants.ConfigEnvDefault)
	}
	viper.SetConfigName(fmt.Sprintf("%s.%s", env, "server"))
	viper.SetConfigType(constants.ConfigEnvFileType)
	viper.AddConfigPath(constants.ConfigEnvFilePath)
	viper.AddConfigPath(constants.ConfigEnvFilePathContainer)
	viper.AddConfigPath(constants.ConfigEnvFilePathContainerConsumer)
	viper.AutomaticEnv()

	var appConf consumer.Configs

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&appConf)
	if err != nil {
		return nil, err
	}

	if strings.HasPrefix(appConf.Consumer.Rabbit.URL, "$") {
		appConf.Consumer.Rabbit.URL = viper.GetString(constants.ConfigAMQPEnvKey)
	}

	return &appConf, nil
}
