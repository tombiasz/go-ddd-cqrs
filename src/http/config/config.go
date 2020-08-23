package config

import "os"

const (
	defaultAppPort = "8000"
)

type AppConfig struct {
	AppPort string
	DbUrl   string
}

func New() AppConfig {
	return AppConfig{
		AppPort: getAppPort(),
		DbUrl:   getDbUrl(),
	}
}

func getAppPort() string {
	p := os.Getenv("APP_PORT")

	if p == "" {
		return defaultAppPort
	}

	return p
}

func getDbUrl() string {
	u := os.Getenv("DATABASE_URL")

	if u == "" {
		panic("DATABASE_URL env var is not set")
	}

	return u
}
