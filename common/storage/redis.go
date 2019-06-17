package storage

import (
	"fmt"
	"github.com/beaconzhang/iris_demo/common"
	"github.com/kataras/iris/sessions/sessiondb/redis"
	"github.com/kataras/iris/sessions/sessiondb/redis/service"
	OriginRedis "github.com/go-redis/redis"
	"sync"
	"time"
)

var (
	redisClient *redis.Database = nil
	redisMutex sync.Mutex

	redisOriginClient *OriginRedis.Client = nil
	redisOriginMutex sync.Mutex
)

func init(){
	GetRedisClient()
	GetOriginRedisClient()
}

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

func GetOriginRedisClient() *OriginRedis.Client{
	if redisOriginClient != nil {
		return redisOriginClient
	}
	redisConf := common.GetConfData().Redis
	redisOriginMutex.Lock()
	defer redisOriginMutex.Unlock()
	if redisOriginClient != nil {
		return redisOriginClient
	}
	redisOriginClient = OriginRedis.NewClient(&OriginRedis.Options{
		Addr:     fmt.Sprintf("%s:%s",redisConf.Host,redisConf.Port),
		Password: redisConf.Passwd,
		DB:       0,
	})
	return redisOriginClient
}

func OriginRedisSet(key string,value string,ttl time.Duration) error{
	return redisOriginClient.Set(makeRedisKey(key),value,ttl).Err()
}

func makeRedisKey(key string)string{
	redisConf := common.GetConfData().Redis
	return redisConf.Prefix + key
}

func OriginRedisGet(key string) (value string,err error){
	ret := redisOriginClient.Get(makeRedisKey(key))
	return ret.Val(),ret.Err()
}