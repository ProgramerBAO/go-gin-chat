package utils

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

// RDB 设置redis全局变量
var RDB *redis.Client

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Network:            "",
		Addr:               viper.GetString("redis.addr"),
		Dialer:             nil,
		OnConnect:          nil,
		Username:           "",
		Password:           viper.GetString("redis.password"),
		DB:                 viper.GetInt("redis.DB"),
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolSize:           viper.GetInt("redis.poolSize"),
		MinIdleConns:       viper.GetInt("redis.minIdleConn"),
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
		Limiter:            nil,
	})
}
