package operation

import (
	"context"
	"gedis"
	"time"
)

// String 处理string
type String struct {
	ctx context.Context
}

// NewString 创建String
func NewString() *String {
	return &String{ctx: context.Background()}
}

// Set 设置string
func (s *String) Set(key string, value interface{}, attrs ...*Attr) *Result {
	expire := Attrs(attrs).Find(AttrExpire).Default(time.Second * 0).(time.Duration)

	nx := Attrs(attrs).Find(AttrNX).Default(nil)
	if nx != nil {
		return NewResult(gedis.Redis().SetNX(s.ctx, key, value, expire).Result())
	}

	xx := Attrs(attrs).Find(AttrXX).Default(nil)
	if xx != nil {
		return NewResult(gedis.Redis().SetXX(s.ctx, key, value, expire).Result())
	}

	return NewResult(gedis.Redis().Set(s.ctx, key, value, expire).Result())
}

// Get redis get
func (s *String) Get(key string) *StringResult {
	return NewStringResult(gedis.Redis().Get(s.ctx, key).Result())
}

// Mget redis mget
func (s *String) Mget(keys ...string) *SliceResult {
	return NewSliceResult(gedis.Redis().MGet(s.ctx, keys...).Result())
}
