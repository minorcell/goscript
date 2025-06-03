package packages

// Filter 过滤一个切片，返回符合条件的元素
// 这是一个简化的例子，使用了 interface{}
// 在 Go 1.18+ 中，可以使用泛型来实现类型安全的过滤
func Filter(slice []interface{}, predicate func(interface{}) bool) []interface{} {
	var result []interface{}
	for _, item := range slice {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// Map 对切片中的每个元素应用一个转换函数
func Map(slice []interface{}, mapper func(interface{}) interface{}) []interface{} {
	result := make([]interface{}, len(slice))
	for i, item := range slice {
		result[i] = mapper(item)
	}
	return result
}

// Reduce 将切片中的元素归约为单个值
func Reduce(slice []interface{}, initialValue interface{}, reducer func(interface{}, interface{}) interface{}) interface{} {
	result := initialValue
	for _, item := range slice {
		result = reducer(result, item)
	}
	return result
}
