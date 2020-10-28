package operation

// StringResult 处理string结果
type StringResult struct {
	Result string
	Err    error
}

// NewStringResult 创建string结果
func NewStringResult(result string, err error) *StringResult {
	return &StringResult{Result: result, Err: err}
}

// Unwrap 返回string结果
func (r *StringResult) Unwrap() string {
	if r.Err != nil {
		panic(r.Err)
	}
	return r.Result
}

// Default 没值时返回string默认值
func (r *StringResult) Default(s string) string {
	if r.Err != nil {
		return s
	}
	return r.Result
}
