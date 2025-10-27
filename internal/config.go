package internal

import (
	"log"
	"os"
	"time"
)

type Config struct {
	Env          string `yaml:"env" env-default:"local"`
	Storage_path string `yaml:"storage_path" end-required:"true"`
	HttpServer
}

type HttpServer struct {
	Address string        `yaml:"address" env-default:"local:9090"`
	Timeout time.Duration `yaml:"timeout" env-default:"4s"`
}

func MustLoad() Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", err)
	}

}
