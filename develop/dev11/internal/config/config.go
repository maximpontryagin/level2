package config

import (
	"os"

	"github.com/joho/godotenv"
)

type ConfigServer struct {
	Host string
	Port string
}

func NewConfigServer() (ConfigServer, error) {
	var cfg ConfigServer
	err := godotenv.Load()
	if err != nil {
		return cfg, err
	}
	cfg.Host = os.Getenv("host")
	cfg.Port = os.Getenv("port")
	return cfg, nil
}
