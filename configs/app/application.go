package app

type Configs struct {
	Application ApplicationConfigs `mapstructure:"application"`
}

type ApplicationConfigs struct {
	Client   ClientConfig   `mapstructure:"client"`
	NewRelic NewRelicConfig `mapstructure:"newrelic"`
	Postgres PGSettings     `mapstructure:"postgres"`
	Server   ServerConfig   `mapstructure:"server"`
}

type ClientConfig struct {
	URL string `mapstructure:"url"`
}

type ServerConfig struct {
	SignKey string `mapstructure:"sign-key"`
	Port    string `mapstructure:"port"`
}

type NewRelicConfig struct {
	ApplicationKey  string `mapstructure:"application-key"`
	ApplicationName string `mapstructure:"application-name"`
}

type PGSettings struct {
	Name     string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Debug    bool   `mapstructure:"debug"`
}
