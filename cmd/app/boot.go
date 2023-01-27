package main

import (
	"fmt"
	"go.uber.org/zap"
	"konntent-workspace-service/configs/app"
	"konntent-workspace-service/pkg/claimer"
	"konntent-workspace-service/pkg/constants"
	"konntent-workspace-service/pkg/nrclient"
	"konntent-workspace-service/pkg/pg"
	pg_migration "konntent-workspace-service/pkg/pg-migration"
	pg_rel_registration "konntent-workspace-service/pkg/pg-rel-registration"
	"time"

	"github.com/spf13/viper"
)

func boot(l *zap.Logger, appConf app.ApplicationConfigs) (*server, error) {
	time.Local, _ = time.LoadLocation("Europe/Istanbul")

	var jwtInstance = claimer.NewClaimer(appConf.Server.SignKey)
	var pgInstance, err = initPG(l, appConf.Postgres)
	if err != nil {
		return nil, err
	}

	//nrInstance, err := initNewRelic(appConf.NewRelic)
	//if err != nil {
	//	return nil, err
	//}

	return &server{
		logger:      l,
		pgInstance:  pgInstance,
		jwtInstance: jwtInstance,
		//nrInstance:  nrInstance,
	}, nil
}

func initPG(l *zap.Logger, cfg app.PGSettings) (pg.Instance, error) {
	return pg.NewPGInstance(l, cfg)
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

	appConf.Application.NewRelic = app.NewRelicConfig{
		ApplicationKey:  viper.GetString(constants.ConfigNRLicenseKey),
		ApplicationName: viper.GetString(constants.ConfigNRAppKey),
	}

	return &appConf, nil
}

func registrar(l *zap.Logger) {
	pg_rel_registration.Register()
}

func migrate(l *zap.Logger, pg pg.Instance) {
	err := pg_migration.Migrate(pg, pg_migration.MigrationModels...)
	if err != nil {
		l.Error("something went wrong while migrating...",
			zap.Error(err))
	}
}
