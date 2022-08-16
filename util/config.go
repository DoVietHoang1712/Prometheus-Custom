package util

type Config struct {
	DBUsername    string `mapstructure:"DB_USERNAME"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	PrometheusUrl string
}

func LoadConfig() (config Config, err error) {
	return Config{
		DBUsername:    "stats",
		DBPassword:    "yW!v6fX6NccJsK",
		DBName:        "stats",
		DBHost:        "localhost",
		DBPort:        "5432",
		PrometheusUrl: "localhost:8428",
	}, nil
}
