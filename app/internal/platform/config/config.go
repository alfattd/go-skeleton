package config

import "os"

type Config struct {
	AppPort        string
	ServiceName    string
	ServiceVersion string
}

func Load() *Config {
	return &Config{
		AppPort:        os.Getenv("APP_PORT"),
		ServiceName:    os.Getenv("SERVICE_NAME"),
		ServiceVersion: os.Getenv("SERVICE_VERSION"),
	}
}
