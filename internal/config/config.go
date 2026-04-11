package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database Database `yaml:"database"`
	Server   Server   `yaml:"server"`
}

type Database struct {
	ConnectionString string `yaml:"connectionString"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func LoadConfig(path string) (*Config, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(content, &config); err != nil {
		return nil, fmt.Errorf("failed to parse yaml config: %w", err)
	}

	if envURL := os.Getenv("DATABASE_URL"); envURL != "" {
		config.Database.ConnectionString = envURL
	}

	if envHost := os.Getenv("SERVER_HOST"); envHost != "" {
		config.Server.Host = envHost
	}

	if envPort := os.Getenv("SERVER_PORT"); envPort != "" {
		config.Server.Port = envPort
	}

	return &config, nil
}
