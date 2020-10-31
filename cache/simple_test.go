package cache

import (
	"gedis/operation"
	"log"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

// TestSimple_Get 测试从缓存获取数据
func TestSimple_Get(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	key := "test:" + strconv.Itoa(rand.Intn(10))
	expire := time.Second * 20

	cache := NewSimple(operation.NewString(), expire)
	cache.Getter = func() string {
		time.Sleep(time.Microsecond * 500)
		log.Println("get data from db") // 模拟从db获取数据
		return "go go go~"
	}

	t.Log(cache.Get(key)) // 第一次没有数据从db获取
	t.Log(cache.Get(key)) // 第二次直接从缓存中获取
}
