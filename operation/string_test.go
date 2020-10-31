package operation

import (
	"testing"
	"time"
)

func TestString_Set(t *testing.T) {
	s := NewString()
	t.Log(s.Set("name", "jj"))
	t.Log(s.Set("tmp_name", "tester", WithExpire(time.Second*20)))
	t.Log(s.Set("now", time.Now().Unix(), WithExpire(time.Second*20), WithNX()))
	t.Log(s.Set("names", "jingjing", WithExpire(time.Second*20), WithXX()))
}

func TestString_Get(t *testing.T) {
	s := NewString()
	t.Log(s.Get("name").Unwrap())
	t.Log(s.Get("xx").Default("oo"))
}

func TestString_Mget(t *testing.T) {
	s := NewString()
	t.Log(s.Mget("name", "age", "gender").Unwrap())

	res := s.Mget("name", "age", "gender").Iter()
	for res.HasNext() {
		t.Log(res.Next())
	}
}
