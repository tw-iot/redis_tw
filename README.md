### redis_tw
redis 操作

go get github.com/tw-iot/redis_tw

### 示例
```
package main

import (
	"fmt"
	"github.com/tw-iot/redis_tw"
	"time"
)

func main() {
	redisCon()
	redisClusterCon()
}

func redisCon()  {
	redisInfo := redis_tw.NewRedisInfo("192.168.146.18:6379", "20201")
	redis_tw.RedisInit(&redisInfo)

	key := "name"
	err := redis_tw.RedisTw.Set(key, "abcd", 10*time.Second).Err()
	if err != nil {
		panic(err)
	}
	val, err := redis_tw.RedisTw.Get(key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(key, ":", val)
}

func redisClusterCon()  {
	redisClusterInfo := redis_tw.NewRedisClusterInfo(
		[]string{"127.0.0.1:7000", "127.0.0.1:7001", "127.0.0.1:7002",
			"127.0.0.1:7003", "127.0.0.1:7004", "127.0.0.1:7005"}, "20201")
	redis_tw.RedisClusterInit(&redisClusterInfo)

	key := "rediskey"
	err := redis_tw.RedisTwCluster.Set(key, "hello", 10*time.Second).Err()
	if err != nil {
		panic(err)
	}
	val, err := redis_tw.RedisTwCluster.Get(key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(key, ":", val)
}

```
