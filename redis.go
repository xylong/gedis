package gedis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"sync"
	"time"
)

var (
	client *redis.Client
	once   sync.Once
)

// Redis 创建redis客户端
func Redis() *redis.Client {
	once.Do(func() {
		client = redis.NewClient(&redis.Options{
			DB:       0,
			Network:  "tcp",
			Addr:     "127.0.0.1:6379",
			Password: "apple",
			// 连接池容量及闲置连接数量
			PoolSize:     15,
			MinIdleConns: 10,
			// 超时
			DialTimeout:  5 * time.Second,
			ReadTimeout:  3 * time.Second,
			WriteTimeout: 3 * time.Second,
			PoolTimeout:  4 * time.Second,
			// 闲置连接检查包括IdleTimeout，MaxConnAge
			IdleCheckFrequency: 60 * time.Second,
			IdleTimeout:        5 * time.Minute,
			MaxConnAge:         0 * time.Second, // 连接存活时长，从创建开始计时，超过指定时长则关闭连接，默认为0，即不关闭存活时长较长的连接
			// 命令执行失败时的重试策略
			MaxRetries:      0,                      // 命令执行失败时，最多重试多少次，默认为0即不重试
			MinRetryBackoff: 8 * time.Millisecond,   //每次计算重试间隔时间的下限，默认8毫秒，-1表示取消间隔
			MaxRetryBackoff: 512 * time.Millisecond, //每次计算重试间隔时间的上限，默认512毫秒，-1表示取消间隔
		})

		pong, err := client.Ping(context.Background()).Result()
		if err != nil {
			log.Fatal(fmt.Errorf("connect error:%s", err))
		}
		log.Println(pong)
	})

	return client
}
