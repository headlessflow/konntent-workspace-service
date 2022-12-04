package constants

type AppEnvironment string

const (
	ConfigEnvDefault                   AppEnvironment = "local"
	ConfigEnvFileType                  string         = "yaml"
	ConfigEnvFilePath                  string         = ".env"
	ConfigEnvFilePathContainer         string         = "/app/.env"
	ConfigEnvFilePathContainerConsumer string         = "/consumer/.env"
)

const (
	ConfigEnvKey       string = "ENV"
	ConfigAMQPEnvKey   string = "AMQP_SERVER_URL"
	ConfigNRAppKey     string = "NEW_RELIC_APP_NAME"
	ConfigNRLicenseKey string = "NEW_RELIC_LICENSE_KEY"
)
