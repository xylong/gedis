package operation

import "testing"

func TestString_Get(t *testing.T) {
	s := NewString()
	t.Log(s.Get("name").Unwrap())
	t.Log(s.Get("xx").Default("oo"))
}
