package config

type Config struct {
	HTTP HTTP
}

type HTTP struct {
	Port int
}

func LoadConfig() (*Config, error) {
	return &Config{
		HTTP: HTTP{
			Port: 8080,
		},
	}, nil
}
