package storage

import (
	"fmt"
	"github.com/beaconzhang/iris_demo/common"
	"github.com/kataras/iris/sessions/sessiondb/redis"
	"github.com/kataras/iris/sessions/sessiondb/redis/service"
	"sync"
	"time"
)

var (
	redisClient *redis.Database = nil
	redisMutex sync.Mutex
)

func GetRedisClient() *redis.Database{
	if redisClient != nil {
		return redisClient
	}
	redisConf := common.GetConfData().Redis
	redisMutex.Lock()
	defer redisMutex.Unlock()
	if redisClient != nil {
		return redisClient
	}
	redisClient = redis.New(service.Config{
		Network: redisConf.Network,
		Addr: fmt.Sprintf("%s:%s",redisConf.Host,redisConf.Port),
		Password: redisConf.Passwd,
		Database: redisConf.Database,
		MaxIdle: redisConf.Maxidle,
		MaxActive: redisConf.Maxactive,
		IdleTimeout: time.Duration(redisConf.Idletimeout)*time.Minute,
		Prefix: redisConf.Prefix,
	})
	return redisClient
}
