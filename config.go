package store

type Config struct{
	URL string `toml: "db_url"`
}

func NewConfig() *Config{
	return &Config{}
}
