package config

import "os"

type Config struct {
	DbUri   string
	AppPort string
}

func Get() *Config {
	conf := Config{}

	conf.DbUri = os.Getenv("DATABASE")
	conf.AppPort = os.Getenv("PORT")

	return &conf
}