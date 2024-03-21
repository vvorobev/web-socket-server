package config

type Config struct {
	HTTP HTTPConfig
}

type HTTPConfig struct {
	Port int
}

func MustLoad() *Config {
	return &Config{
		HTTP: HTTPConfig{Port: 3003},
	}
}
