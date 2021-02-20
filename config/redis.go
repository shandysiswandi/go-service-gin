package config

import (
	"os"
	"strconv"
)

// RedisConfig is
type RedisConfig struct {
	Address  string
	Password string
	Database int
}

// NewRedisConfig is
func NewRedisConfig() *RedisConfig {
	di, err := strconv.Atoi(os.Getenv("REDIS_DATABASE"))
	if err != nil {
		di = 0
	}
	return &RedisConfig{
		Address:  os.Getenv("REDIS_ADDRESS"),
		Password: os.Getenv("REDIS_PASSWORD"),
		Database: di,
	}
}
