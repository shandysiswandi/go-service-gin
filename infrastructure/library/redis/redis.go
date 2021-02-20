package redis

import (
	"go-service-gin/config"

	lib "github.com/go-redis/redis"
)

// Redis is service for caching
type Redis struct {
	client *lib.Client
}

// New is constructor
func New(rc *config.RedisConfig) *Redis {
	if rc == nil {
		return nil
	}

	return &Redis{lib.NewClient(&lib.Options{
		Addr:     rc.Address,
		Password: rc.Password,
		DB:       rc.Database,
	})}
}

// Get is
func (Redis) Get() error {
	return nil
}

// Set is
func (Redis) Set() error {
	return nil
}

// Increment is
func (Redis) Increment() error {
	return nil
}

// Decrement is
func (Redis) Decrement() error {
	return nil
}
