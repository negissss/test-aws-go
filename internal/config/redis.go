package config

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

type RedisConfig struct {
	Addr    string
	Auth    string
	Cluster bool
}

func LoadRedisConfig() *RedisConfig {
	clusterStr := os.Getenv("CLUSTER")
	cluster, err := strconv.ParseBool(clusterStr)
	if err != nil {
		logrus.Printf("Invalid value for CLUSTER: %s, defaulting to false", clusterStr)
		cluster = false
	}

	return &RedisConfig{
		Addr:    os.Getenv("REDIS_HOST"),
		Auth:    os.Getenv("REDIS_AUTH"),
		Cluster: cluster,
	}
}
