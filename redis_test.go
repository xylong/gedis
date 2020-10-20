package gedis

import (
	"context"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	ctx := context.Background()
	redis := Redis(0, "127.0.0.1:6379", "apple")
	expire := time.Second * 3
	expired := expire + 1

	redis.Set(ctx, "name", "golang", expire)
	if v, err := redis.Get(ctx, "name").Result(); err != nil {
		t.Error(err.Error())
	} else {
		t.Log(v)
	}

	time.Sleep(expired)
	if v, err := redis.Get(ctx, "name").Result(); err != nil && err.Error() == "redis: nil" {
		t.Log(v)
	}
}
