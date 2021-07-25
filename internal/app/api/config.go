package api

import "github.com/srm-developer/standardWebServer/storage"

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LoggerLevel string `toml:"logger_level"`
	Storage     *storage.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8084",
		LoggerLevel: "debug",
		Storage:     storage.NewConfig(),
	}
}
