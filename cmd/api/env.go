package main

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	Env         string `default:"dev"`
	APIPort     int    `default:3000`
	LogLevel    string `default:"debug"`
	DMLInitFile string
}

func load() (cfg config, err error) {
	_ = godotenv.Load("cmd/api/.env")
	err = envconfig.Process("", &cfg)
	return
}
