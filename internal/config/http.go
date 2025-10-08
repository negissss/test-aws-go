package config

import "os"

type HTTPConfig struct {
	Host       string
	Port       string
	ExposePort string
	AppEnv     string
}

func LoadHTTPConfig() *HTTPConfig {
	return &HTTPConfig{
		Host:       getEnv("HOST", "0.0.0.0"),
		Port:       getEnv("PORT", "8080"),
		ExposePort: getEnv("EXPOSE_PORT", "8080"),
		AppEnv:     getEnv("GIN_MODE", "release"),
	}
}

func (c *HTTPConfig) IsDevelopment() bool {
	return c.AppEnv == "debug"
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
