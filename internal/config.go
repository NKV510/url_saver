package internal

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Env          string `yaml:"env"`
	Storage_path string `yaml:"storage_path"`
	Database
	HTTP_server
}

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	Ssl_mode string `yaml:"ssl_mode"`
}

type HTTP_server struct {
	Address string        `yaml:"address"`
	Timeout time.Duration `yaml:"timeout"`
}

func Load() (*Config, error) {
	data, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		return nil, fmt.Errorf("can not read config: %w", err)
	}
	var cfg Config
	if err = yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("can not parsing yaml file: %w", err)
	}
	return &cfg, nil
}
func (c *Config) DatabaseURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		c.Database.User,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Name,
		c.Database.Ssl_mode)
}
