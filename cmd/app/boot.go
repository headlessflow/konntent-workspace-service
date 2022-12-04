package main

import (
	"fmt"
	"konntent-workspace-service/configs/app"
	"konntent-workspace-service/pkg/claimer"
	"konntent-workspace-service/pkg/constants"
	"konntent-workspace-service/pkg/dummyclient"
	"konntent-workspace-service/pkg/httpclient"
	"konntent-workspace-service/pkg/nrclient"
	"konntent-workspace-service/pkg/rabbit"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func boot(l *logrus.Logger, appConf app.ApplicationConfigs) (*server, error) {
	time.Local, _ = time.LoadLocation("Europe/Istanbul")

	var (
		httpClient         = httpclient.NewHTTPClient()
		dummyClient        = dummyclient.NewClient(initDummyClientConfig(appConf.Client), httpClient)
		mqConnectionClient = rabbit.NewClientConnector(l)
		mqPreProducer      = rabbit.NewPreProducer(l, appConf.Rabbit.QueueSettings)
		mqProducer         = rabbit.NewMessagingClient(l, mqConnectionClient, mqPreProducer)
		jwtInstance        = claimer.NewClaimer(appConf.Server.SignKey)
	)

	//err := mqProducer.ConnectToBroker(appConf.Rabbit.URL)
	//if err != nil {
	//	return nil, err
	//}

	//nrInstance, err := initNewRelic(appConf.NewRelic)
	//if err != nil {
	//	return nil, err
	//}

	return &server{
		logger:      l,
		dummyClient: dummyClient,
		mqProducer:  mqProducer,
		jwtInstance: jwtInstance,
		//nrInstance:  nrInstance,
	}, nil
}

func initDummyClientConfig(dummyConf app.ClientConfig) dummyclient.Config {
	return dummyclient.Config{
		URL: dummyConf.URL,
	}
}

func initNewRelic(cfg app.NewRelicConfig) (nrclient.NewRelicInstance, error) {
	return nrclient.InitNewRelic(nrclient.Config{
		Key:     cfg.ApplicationKey,
		AppName: cfg.ApplicationName,
	})
}

func initConfig[T string | constants.AppEnvironment](env T) (*app.Configs, error) {
	if env == "" {
		env = T(constants.ConfigEnvDefault)
	}
	viper.AutomaticEnv()
	viper.SetConfigName(fmt.Sprintf("%s.%s", env, "server"))
	viper.SetConfigType(constants.ConfigEnvFileType)
	viper.AddConfigPath(constants.ConfigEnvFilePath)
	viper.AddConfigPath(constants.ConfigEnvFilePathContainer)

	var appConf app.Configs

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&appConf)
	if err != nil {
		return nil, err
	}

	if strings.HasPrefix(appConf.Application.Rabbit.URL, "$") {
		appConf.Application.Rabbit.URL = viper.GetString(constants.ConfigAMQPEnvKey)
	}

	appConf.Application.NewRelic = app.NewRelicConfig{
		ApplicationKey:  viper.GetString(constants.ConfigNRLicenseKey),
		ApplicationName: viper.GetString(constants.ConfigNRAppKey),
	}

	return &appConf, nil
}
