package app

type Configs struct {
	Application ApplicationConfigs `mapstructure:"application"`
}

type ApplicationConfigs struct {
	Client   ClientConfig   `mapstructure:"client"`
	Rabbit   RabbitConfig   `mapstructure:"rabbitmq"`
	NewRelic NewRelicConfig `mapstructure:"newrelic"`
	Server   ServerConfig   `mapstructure:"server"`
}

type ClientConfig struct {
	URL string `mapstructure:"url"`
}

type RabbitConfig struct {
	URL           string              `mapstructure:"server-url"`
	QueueSettings RabbitQueueSettings `mapstructure:"queue-settings"`
}

type ServerConfig struct {
	SignKey string `mapstructure:"sign-key"`
}

type NewRelicConfig struct {
	ApplicationKey  string `mapstructure:"application-key"`
	ApplicationName string `mapstructure:"application-name"`
}

type RabbitQueueSettings struct {
	ExchangeName string `mapstructure:"exchange"`
	QueueName    string `mapstructure:"queue"`
	RoutingKey   string `mapstructure:"routing"`
	RoutingTTL   int64  `mapstructure:"ttl"`
	Dlx          string `mapstructure:"dlx"`
	Dlq          string `mapstructure:"dlq"`
	Dlrk         string `mapstructure:"dlrk"`
	DlqTTL       int64  `mapstructure:"dlq-ttl"`
	MaxRetries   int    `mapstructure:"max-retries"`
}
