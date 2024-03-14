package api

import "github.com/dobrbi2000/web/server_db2_handler/storage"

//main instance

type Config struct {
	//port
	BindAddr string `toml:"bind_addr"`
	//Logger Level
	LoggerLevel string `toml:"logger_level"`
	//Store configs
	Storage *storage.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8080",
		LoggerLevel: "debug",
		Storage:     storage.NewConfig(),
	}
}
