package cache

import (
	"github.com/xylong/gedis/operation"
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

	cache := NewSimple(operation.NewString(), expire, SerializeJson, nil)
	cache.Getter = func() interface{} {
		time.Sleep(time.Microsecond * 500)
		log.Println("get data from db") // 模拟从db获取数据
		return "go go go~"
	}

	t.Log(cache.Get(key)) // 第一次没有数据从db获取
	t.Log(cache.Get(key)) // 第二次直接从缓存中获取
}

// TestSimple_Get2 测试策略
func TestSimple_Get2(t *testing.T) {
	expire := time.Second * 20
	cache := NewSimple(operation.NewString(), expire, SerializeJson, operation.NewPenetratePolicy("^test:([1-9]|[1-9]\\d|100)$"))
	cache.Getter = func() interface{} {
		return time.Now().Unix()
	}

	keys := []string{"test:a", "test:-1", "test:0", "test:101", "test:100.5", "test:100"}
	c := make(chan string)
	for _, key := range keys {
		go func(key string) {
			defer func() {
				if err := recover(); err != nil {
					t.Log(err)
				}
			}()

			r := cache.Get(key)
			t.Log(key, r)
			c <- key
		}(key)
	}
	if "test:100" == <-c {
		t.Log("ok")
	} else {
		t.Error("failed")
	}
}
