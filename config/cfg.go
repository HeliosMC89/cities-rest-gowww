package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
	SslMode  string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "postgres",
			Host:     "localhost",
			Port:     5432,
			Username: "helios",
			Password: "lfc1892",
			Name:     "go-web",
			Charset:  "utf8",
			SslMode:  "disable",
		},
	}
}
