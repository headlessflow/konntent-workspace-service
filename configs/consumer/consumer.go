package consumer

type Configs struct {
	Consumer Application `mapstructure:"application"`
}

type Application struct {
	Client ClientConfig `mapstructure:"client"`
	Rabbit RabbitConfig `mapstructure:"rabbitmq"`
}

type ClientConfig struct {
	URL string `mapstructure:"url"`
}

type RabbitConfig struct {
	URL           string              `mapstructure:"server-url"`
	QueueSettings RabbitQueueSettings `mapstructure:"queue-settings"`
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
