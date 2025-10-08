package config

import "os"

type BtcConfig struct {
	BtcRpcURL   string
	BtcRpcToken string
}

func LoadBtcConfig() *BtcConfig {
	return &BtcConfig{
		BtcRpcURL:   os.Getenv("BTC_RPC_URL"),
		BtcRpcToken: os.Getenv("BTC_RPC_TOKEN"),
	}
}
