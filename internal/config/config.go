package config

type Config struct {
	DB   *DBConfig
	HTTP *HTTPConfig
	// BTC   *BtcConfig
	// Redis *RedisConfig
	// EVM   *EvmConfig
	// CMC   *CmcConfig
}

func NewConfig() *Config {
	return &Config{
		// DB:   LoadDBConfig(),
		HTTP: LoadHTTPConfig(),
		// BTC:   LoadBtcConfig(),
		// Redis: LoadRedisConfig(),
		// EVM:   LoadEvmConfig(),
		// CMC:   LoadCmcConfig(),
	}
}
