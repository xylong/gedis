package example

import (
	"encoding/json"
	"fmt"
	"gedis/cache"
	"gedis/operation"
	"log"
	"math/rand"
	"sync"
	"time"
)

// UserCachePool 用户缓存池
var UserCachePool *sync.Pool

func init() {
	UserCachePool = &sync.Pool{New: func() interface{} {
		return cache.NewSimple(operation.NewString(), time.Minute*1)
	}}
}

// UserCache 从缓存池获取用户信息
func UserCache() *cache.Simple {
	return UserCachePool.Get().(*cache.Simple)
}

// ReleaseUser 放回缓存池
func ReleaseUser(simple *cache.Simple) {
	UserCachePool.Put(simple)
}

// userGetter 获取用户信息
// 模拟从数据库获取数据
func userGetter(id int) cache.DBGetterFunc {
	return func() string {
		user := struct {
			ID   int
			Name string
			Age  int
		}{
			ID:   id,
			Name: "静静",
			Age:  18,
		}
		if bytes, err := json.Marshal(&user); err != nil {
			log.Fatal(err)
			return ""
		} else {
			return string(bytes)
		}
	}
}

// UserExample 获取用户信息
func UserExample() string {
	uc := UserCache()
	defer ReleaseUser(uc)

	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(100)
	uc.Getter = userGetter(id)
	return uc.Get(fmt.Sprintf("users:%d", id)).(string)
}
