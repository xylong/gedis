package example

import (
	"fmt"
	"github.com/xylong/gedis/cache"
	"github.com/xylong/gedis/operation"
	"math/rand"
	"sync"
	"time"
)

var (
	// UserCacheJsonPool json缓存池
	UserCacheJsonPool *sync.Pool
	// UserCacheGobPool gob缓存池
	UserCacheGobPool *sync.Pool
)

func init() {
	UserCacheJsonPool = &sync.Pool{New: func() interface{} {
		return cache.NewSimple(
			operation.NewString(), time.Minute*1, cache.SerializeJson, operation.NewPenetratePolicy("^users:\\d{1,100}$"))
	}}
	UserCacheGobPool = &sync.Pool{New: func() interface{} {
		return cache.NewSimple(
			operation.NewString(), time.Minute*1, cache.SerializeGob, operation.NewPenetratePolicy("^users:\\d{1,100}$"))
	}}
}

// UserCache 从缓存池获取用户信息
func UserCache() *cache.Simple {
	return UserCacheJsonPool.Get().(*cache.Simple)
}

func UserCache2() *cache.Simple {
	return UserCacheGobPool.Get().(*cache.Simple)
}

// ReleaseUser 放回缓存池
func ReleaseUser(simple *cache.Simple) {
	UserCacheJsonPool.Put(simple)
}

func ReleaseUser2(simple *cache.Simple) {
	UserCacheGobPool.Put(simple)
}

// User 用户
type User struct {
	ID   int
	Name string
	Age  int
}

// userGetter 获取用户信息
// 模拟从数据库获取数据
func userGetter(id int) cache.DBGetterFunc {
	return func() interface{} {
		user := &User{
			ID:   id,
			Name: "静静",
			Age:  18,
		}
		return &user
	}
}

// UserJsonExample 获取用户信息
func UserJsonExample() string {
	// 1.从对象池获取用户缓存
	uc := UserCache()
	defer ReleaseUser(uc)
	// 2.获取参数，设置getter
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(100)
	uc.Getter = userGetter(id)
	// 3.获取缓存(如果没有缓存，则上面getter调用)
	return uc.Get(fmt.Sprintf("users:%d", id)).(string)
}

// UserGobExample 获取用户信息
func UserGobExample() *User {
	// 1.从对象池获取用户缓存
	uc := UserCache2()
	defer ReleaseUser2(uc)
	// 2.获取参数，设置getter
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(100)
	uc.Getter = userGetter(id)
	// 3.获取缓存(如果没有缓存，则上面getter调用)
	user := &User{}
	return uc.GetObject(fmt.Sprintf("users:%d", id), user).(*User)
}
