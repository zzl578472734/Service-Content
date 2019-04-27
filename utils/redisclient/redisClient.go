package redisclient

import (
	"encoding/json"
	"errors"
	"github.com/go-redis/redis"
	"time"
)

const DefaultCacheExpire = 10 * time.Minute

var (
	RedisClient *redis.Client
)

type Config struct {
	Host        string
	Password    string
	Port        int
	Database    int
	DialTimeout int
}

func InitRedis(redisMode int, config *Config) {
	switch redisMode {
	case 1:
		initStandAlone(config)
		RedisClient = RedisStandAloneClient
	case 2:
		InitSentinel()
		RedisClient = SentinelClient
	case 3:
		InitCluster()
	}
}

func SetCache(key string, value interface{}, expiration ...time.Duration) error {
	if key == "" ||
		len(key) <= 0 {
		return errors.New("缓存的key不存在")
	}

	if value == nil {
		return errors.New("缓存的value不存在")
	}

	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	if len(bytes) <= 0 {
		return nil
	}

	expire := DefaultCacheExpire
	if len(expiration) > 0 {
		expire = expiration[0]
	}

	err = RedisClient.Set(key, bytes, expire).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetCache(key string, value interface{}) error {
	if key == "" ||
		len(key) <= 0 {
		return errors.New("缓存的key不存在")
	}

	if value == nil {
		return errors.New("缓存的value不存在")
	}

	bytes, err := RedisClient.Get(key).Bytes()
	if RedisErr(err) != nil {
		return err
	}

	if len(bytes) <= 0 {
		return nil
	}

	err = json.Unmarshal(bytes, value)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCache(keys ...string) error {
	if len(keys) <= 0 {
		return errors.New("缓存的key不存在")
	}

	err := RedisClient.Del(keys...).Err()
	if err != nil {
		return err
	}
	return nil
}

func RedisErr(err error) error {
	if err != redis.Nil {
		return err
	}

	return nil
}
