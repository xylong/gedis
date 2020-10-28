package operation

// SliceResult 处理slice结果
type SliceResult struct {
	Result []interface{}
	Err    error
}

// NewSliceResult 创建slice结果
func NewSliceResult(result []interface{}, err error) *SliceResult {
	return &SliceResult{Result: result, Err: err}
}

// Unwrap 返回slice结果
func (r *SliceResult) Unwrap() []interface{} {
	if r.Err != nil {
		panic(r.Err)
	}
	return r.Result
}

// Default 没值时返回slice默认值
func (r *SliceResult) Default(any []interface{}) []interface{} {
	if r.Err != nil {
		return any
	}
	return r.Result
}
