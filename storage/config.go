package storage

type Config struct {
	//подключение к БД
	DatabaseURL string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{}
}
