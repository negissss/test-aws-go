package config

import "os"

type CmcConfig struct {
	CmcApiKey string
}

func LoadCmcConfig() *CmcConfig {
	return &CmcConfig{
		CmcApiKey: os.Getenv("CMC_API_KEY"),
	}
}
