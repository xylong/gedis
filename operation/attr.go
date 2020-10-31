package operation

import (
	"fmt"
	"time"
)

const (
	// AttrExpire 过期时间
	AttrExpire = "expire"
	// AttrNotFound 属性未找到
	AttrNotFound = "operation attr:%s not found error"
	// AttrNx 表示SetNX
	AttrNX = "nx"
	// AttrXx 表示SetXX
	AttrXX = "xx"
)

// empty 空
type empty struct{}

// Attr 属性
type Attr struct {
	Name  string
	Value interface{}
}

// Attrs 属性切片
type Attrs []*Attr

// Find 查找属性
func (a Attrs) Find(name string) *Result {
	for _, attr := range a {
		if attr.Name == name {
			return NewResult(attr.Value, nil)
		}
	}
	return NewResult(nil, fmt.Errorf(AttrNotFound, name))
}

// WithExpire 设置过期时间属性
func WithExpire(t time.Duration) *Attr {
	return &Attr{
		Name:  AttrExpire,
		Value: t,
	}
}

// WithExpire 设置SetNX属性
// key不存在时设置
func WithNX() *Attr {
	return &Attr{
		Name:  AttrNX,
		Value: empty{},
	}
}

// WithXX 设置SetXX属性
// key存在时设置
func WithXX() *Attr {
	return &Attr{
		Name:  AttrXX,
		Value: empty{},
	}
}
