package utils

import (
	"github.com/go-redis/redis"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego"
	"fmt"
	"time"
)

var (
	RedisClient *redis.Client
)

func InitRedis() {
	redisMode,err := beego.AppConfig.Int("redisMode")
	if err != nil{
		panic(err)
	}
	switch redisMode {
	case 0:
		initRedisStandAlone()
	case 1:
		initRedisSentinel()
	case 2:
		initRedisCluster()
	}
}

func initRedisStandAlone() {
	logs.Info("start init redis service....")
	redisHost := beego.AppConfig.DefaultString("redis.host", "")
	if len(redisHost) <= 0 {
		panic("redis host is empty")
	}

	redisPort, err := beego.AppConfig.Int("redis.port")
	if err != nil {
		panic(err)
	}

	redisDatabase, err := beego.AppConfig.Int("redis.database")
	if err != nil {
		panic(err)
	}

	redisPassword := beego.AppConfig.DefaultString("redis.password", "")

	redisDialTimeout, err := beego.AppConfig.Int("redisDialTimeout")
	if err != nil {
		panic(redisDialTimeout)
	}

	redisAddr := fmt.Sprintf("%s:%d", redisHost, redisPort)

	logs.Info("redis connect adddr: %s, database: %d, password:%s, dialTimeout:%v", redisAddr, redisDatabase, redisPassword, redisDialTimeout)
	RedisClient = redis.NewClient(&redis.Options{
		Addr:        redisAddr,
		DB:          redisDatabase,
		Password:    redisPassword,
		DialTimeout: time.Duration(redisDialTimeout) * time.Second,
	})

	if _, err := RedisClient.Ping().Result(); err != nil {
		panic(err)
	}
	logs.Info("finish initial redis...")
}

func initRedisSentinel()  {
	
}

func initRedisCluster()  {
	
}

func RedisErr(err error) error {
	if err != redis.Nil {
		return err
	}

	return nil
}
