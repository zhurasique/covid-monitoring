package apiserver

import "covid-monitoring/internal/app/store"

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8084",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
