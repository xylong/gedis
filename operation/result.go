package operation

// Result 任意类型返回结果
type Result struct {
	// Value 返回值
	Value interface{}
	// Err 错误
	Err error
}

// NewResult 创建结果
func NewResult(result interface{}, err error) *Result {
	return &Result{
		Value: result,
		Err:   err,
	}
}

// Unwrap 返回结果
func (r *Result) Unwrap() interface{} {
	if r.Err != nil {
		panic(r.Err)
	}
	return r.Value
}

// Default 没值时返回默认值
func (r *Result) Default(value interface{}) interface{} {
	if r.Err != nil {
		return value
	}
	return r.Value
}
