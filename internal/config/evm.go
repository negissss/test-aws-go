package config

import "os"

type EvmConfig struct {
	EthRpcURL string
}

func LoadEvmConfig() *EvmConfig {
	return &EvmConfig{
		EthRpcURL: os.Getenv("ETH_RPC_URL"),
	}
}
