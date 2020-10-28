package operation

import (
	"context"
	"gedis"
)

// String 处理string
type String struct {
	ctx context.Context
}

// NewString 创建String
func NewString() *String {
	return &String{ctx: context.Background()}
}

func (s *String) Set() {

}

func (s *String) Get(key string) *StringResult {
	return NewStringResult(gedis.Redis().Get(s.ctx, key).Result())
}
