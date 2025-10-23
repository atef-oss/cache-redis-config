package helpers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-redsync/redsync"
)

func GetRedisLockKey(namespace string, appID string) string {
	return fmt.Sprintf("%s_%s_lock", namespace, appID)
}

func GetRedisKey(namespace string) string {
	return namespace
}

func GetRedisClient(rds *redsync.Redsync) (*redsync.Pool, error) {
	if rds == nil {
		return nil, fmt.Errorf("Redis instance is not set")
	}
	return rds.GetPool(), nil
}

func GetRedisClientFromConfig(redisConfig RedisConfig) (*redsync.Pool, error) {
	if redisConfig == nil {
		return nil, fmt.Errorf("Redis configuration is not set")
	}
	if redisConfig.Address == "" {
		return nil, fmt.Errorf("Redis address is not set")
	}
	if redisConfig.PoolSize == 0 {
		return nil, fmt.Errorf("Redis pool size is not set")
	}
	p := redsync.NewPool(&redsync.Config{
		SentinelAddrs: []string{redisConfig.Address},
		Password:      redisConfig.Password,
		DB:           redisConfig.DB,
		SentinelPassword: func(s string) (string, error) {
			return redisConfig.Password, nil
		},
		PoolSize: redisConfig.PoolSize,
	})
	return p, nil
}

type RedisConfig struct {
	Address    string
	Password   string
	DB         int
	PoolSize   int
	SentinelID string
}

func ParseRedisConfigFromEnv() (RedisConfig, error) {
	address := os.Getenv("REDIS_ADDRESS")
	if address == "" {
		return RedisConfig{}, fmt.Errorf("REDIS_ADDRESS is not set")
	}
	password := os.Getenv("REDIS_PASSWORD")
	if password == "" {
		return RedisConfig{}, fmt.Errorf("REDIS_PASSWORD is not set")
	}
	db := os.Getenv("REDIS_DB")
	if db == "" {
		db = "0"
	}
	db, err := strconv.Atoi(db)
	if err != nil {
		return RedisConfig{}, fmt.Errorf("REDIS_DB is not a valid integer")
	}
	poolSize := os.Getenv("REDIS_POOL_SIZE")
	if poolSize == "" {
		poolSize = "10"
	}
	poolSize, err := strconv.Atoi(poolSize)
	if err != nil {
		return RedisConfig{}, fmt.Errorf("REDIS_POOL_SIZE is not a valid integer")
	}
	sentinelID := os.Getenv("REDIS_SENTINEL_ID")
	if sentinelID == "" {
		sentinelID = "mymaster"
	}
	return RedisConfig{
		Address:    address,
		Password:   password,
		DB:         db,
		PoolSize:   poolSize,
		SentinelID: sentinelID,
	}, nil
}