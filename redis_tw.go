package redis_tw

import (
	"github.com/go-redis/redis"
)

var RedisTw *redis.Client
var RedisTwCluster *redis.ClusterClient

type RedisInfo struct {
	Addr         string
	Password     string
	Db           int
	PoolSize     int
	MinIdleConns int
}

type RedisClusterInfo struct {
	Addrs        []string
	Password     string
	PoolSize     int
	MinIdleConns int
}

func NewRedisInfo(addr, password string) RedisInfo {
	return RedisInfo{
		Addr:         addr,
		Password:     password,
		Db:           0,
		PoolSize:     0,
		MinIdleConns: 0,
	}
}

func NewRedisClusterInfo(addrs []string, password string) RedisClusterInfo {
	return RedisClusterInfo{
		Addrs:        addrs,
		Password:     password,
		PoolSize:     0,
		MinIdleConns: 0,
	}
}

func RedisInit(redisInfo *RedisInfo) *redis.Client {
	RedisTw = redis.NewClient(&redis.Options{
		Addr:         redisInfo.Addr,
		Password:     redisInfo.Password, // no password set
		DB:           redisInfo.Db,       // use default DB
		PoolSize:     redisInfo.PoolSize,
		MinIdleConns: redisInfo.MinIdleConns,
	})
	return RedisTw
}

func RedisClose() {
	RedisTw.Close()
}

func RedisClusterInit(redisClusterInfo *RedisClusterInfo) *redis.ClusterClient {
	RedisTwCluster := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        redisClusterInfo.Addrs,
		Password:     redisClusterInfo.Password,
		PoolSize:     redisClusterInfo.PoolSize,
		MinIdleConns: redisClusterInfo.MinIdleConns,
	})
	return RedisTwCluster
}

func RedisClusterClose() {
	RedisTwCluster.Close()
}
