package cache

import (
	"gedis/operation"
	"time"
)

// DBGetterFunc db操作函数
type DBGetterFunc func() string

// Simple 简单缓存
type Simple struct {
	// Operation redis操作
	Operation *operation.String
	// Expire 过期时间
	Expire time.Duration
	// Getter 数据库获取数据的函数
	Getter DBGetterFunc
}

// NewSimple 创建简单缓存
func NewSimple(operation *operation.String, expire time.Duration) *Simple {
	return &Simple{Operation: operation, Expire: expire}
}

// Set 设置缓存
func (s *Simple) Set(key string, value interface{}) {
	s.Operation.Set(key, value, operation.WithExpire(s.Expire)).Unwrap()
}

// Get 获取缓存
func (s *Simple) Get(key string) (value interface{}) {
	value = s.Operation.Get(key).UnwrapElse(s.Getter) // 如果缓存未获取到则从数据库获取
	s.Set(key, value)
	return
}
