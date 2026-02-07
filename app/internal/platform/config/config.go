package config

import "os"

type Config struct {
	ServiceName    string
	ServiceVersion string
}

func Load() *Config {
	name := os.Getenv("SERVICE_NAME")
	version := os.Getenv("SERVICE_VERSION")

	return &Config{
		ServiceName:    name,
		ServiceVersion: version,
	}
}
