package operation

import (
	"testing"
)

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
