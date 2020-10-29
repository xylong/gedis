package operation

// Iterator 迭代器
type Iterator struct {
	data  []interface{}
	index int
}

// NewIterator 创建迭代器
func NewIterator(data []interface{}) *Iterator {
	return &Iterator{data: data}
}

// HasNext 判断是否继续迭代
func (i *Iterator) HasNext() bool {
	length := len(i.data)
	if i.data == nil || length == 0 {
		return false
	}
	return i.index < length
}

// Next 取slice中的下一个值
func (i *Iterator) Next() (any interface{}) {
	any = i.data[i.index]
	i.index++
	return
}
