package redisclient

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"time"
)

var (
	RedisStandAloneClient *redis.Client
)

func initStandAlone(config *Config) {
	log.Printf("start init redis service....")
	if len(config.Host) <= 0 {
		panic("redis host is err")
	}

	if config.Port <= 0 {
		panic("redis port is err")
	}

	if config.Database < 0 {
		panic("redis database is err")
	}

	redisAddr := fmt.Sprintf("%s:%d", config.Host, config.Port)

	log.Printf("redis connect adddr: %s, database: %d, password:%s, dialTimeout:%v", redisAddr, config.Database, config.Password, config.DialTimeout)
	RedisStandAloneClient = redis.NewClient(&redis.Options{
		Addr:        redisAddr,
		DB:          config.Database,
		Password:    config.Password,
		DialTimeout: time.Duration(config.DialTimeout) * time.Second,
	})

	if _, err := RedisStandAloneClient.Ping().Result(); err != nil {
		panic(err)
	}
	log.Printf("finish initial redis...")
}
