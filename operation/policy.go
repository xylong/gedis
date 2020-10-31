package operation

import (
	"fmt"
	"regexp"
)

const (
	PolicyMsg = "error cache key: %s"
)

// Policy 缓存策略
type Policy interface {
	Before(key string)
}

// PenetratePolicy 缓存穿透策略
type PenetratePolicy struct {
	KeyRegx string // 检查key的正则
}

// NewPenetratePolicy 创建策略
func NewPenetratePolicy(keyRegx string) *PenetratePolicy {
	return &PenetratePolicy{KeyRegx: keyRegx}
}

func (p *PenetratePolicy) Before(key string) {
	if !regexp.MustCompile(p.KeyRegx).MatchString(key) {
		panic(fmt.Sprintf(PolicyMsg, key))
	}
}
