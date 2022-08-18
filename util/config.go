package util

import (
	"regexp"
	"strconv"
)

type Config struct {
	DBUsername    string `mapstructure:"DB_USERNAME"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	PrometheusUrl string
	TimeStart     string
	TimeZone      string
}

func LoadConfig() (config Config, err error) {
	return Config{
		DBUsername:    "stats",
		DBPassword:    "yW!v6fX6NccJsK",
		DBName:        "stats",
		DBHost:        "localhost",
		DBPort:        "5432",
		PrometheusUrl: "localhost:8428",
		TimeStart:     "15:07",
		TimeZone:      "Asia/Bangkok",
	}, nil
}

func GetTimeStart(timeStart string) (hour, minute int64) {
	regex := regexp.MustCompile("(?P<hour>\\d+):(?P<minute>\\d+)")
	check := regex.Match([]byte(timeStart))
	if check {
		res := regex.FindStringSubmatch(timeStart)
		hour, _ = strconv.ParseInt(res[1], 10, 32)
		minute, _ = strconv.ParseInt(res[2], 10, 32)
	}
	return hour, minute
}
