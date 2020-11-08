package cache

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"github.com/xylong/gedis/operation"
	"log"
	"reflect"
	"time"
)

const (
	// SerializeJson json序列化
	SerializeJson = "json"
	// SerializeGob gob序列化
	SerializeGob = "gob"
)

// DBGetterFunc db操作函数
type DBGetterFunc func() interface{}

// Simple 简单缓存
type Simple struct {
	// Operation redis操作
	Operation *operation.String
	// Expire 过期时间
	Expire time.Duration
	// Getter 数据库获取数据的函数
	Getter DBGetterFunc
	// Serialization 序列化方式
	Serialization string
	// Policy 策略
	Policy operation.Policy
}

// NewSimple 创建简单缓存
func NewSimple(operation *operation.String, expire time.Duration, serialization string, policy operation.Policy) *Simple {
	return &Simple{Operation: operation, Expire: expire, Serialization: serialization, Policy: policy}
}

// Set 设置缓存
func (s *Simple) Set(key string, value interface{}) {
	s.Operation.Set(key, value, operation.WithExpire(s.Expire)).Unwrap()
}

// Get 获取缓存
func (s *Simple) Get(key string) (value interface{}) {
	// 检查策略
	if s.Policy != nil {
		s.Policy.Before(key)
	}

	var f func() string
	obj := s.Getter()

	switch s.Serialization {
	case SerializeJson:
		f = func() string {
			if b, err := json.Marshal(obj); err != nil {
				log.Fatal(err)
				return ""
			} else {
				return string(b)
			}
		}
	case SerializeGob:
		f = func() string {
			buff := &bytes.Buffer{}
			encode := gob.NewEncoder(buff)
			if err := encode.Encode(obj); err != nil {
				return ""
			}
			return buff.String()
		}
	}

	value = s.Operation.Get(key).UnwrapElse(f)
	s.Set(key, value)
	return
}

// GetObject 获取缓存对象
func (s *Simple) GetObject(key string, obj interface{}) interface{} {
	// 判断obj是否是指针，否则obj无法赋值
	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Ptr {
		return nil
	}

	result := s.Get(key)
	if result == nil {
		return nil
	}

	switch s.Serialization {
	case SerializeJson:
		if err := json.Unmarshal([]byte(result.(string)), obj); err != nil {
			return nil
		}
	case SerializeGob:
		buff := &bytes.Buffer{}
		buff.WriteString(result.(string))
		decode := gob.NewDecoder(buff)
		if decode.Decode(obj) != nil {
			return nil
		}
	}

	return obj
}
