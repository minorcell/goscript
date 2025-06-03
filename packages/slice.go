package packages

import "sort"

// Filter 过滤一个切片，返回符合条件的元素
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

// Max 返回切片中的最大值
func Max(slice []interface{}) interface{} {
	max := slice[0]
	for _, item := range slice {
		if item.(float64) > max.(float64) {
			max = item
		}
	}
	return max
}

// Min 返回切片中的最小值
func Min(slice []interface{}) interface{} {
	min := slice[0]
	for _, item := range slice {
		if item.(float64) < min.(float64) {
			min = item
		}
	}
	return min
}

// Splice 替换切片中的元素
func Splice(slice []interface{}, start int, deleteCount int, items ...interface{}) []interface{} {
	result := make([]interface{}, len(slice))
	copy(result, slice[:start])
	copy(result[start:], items)
	copy(result[start+len(items):], slice[start+deleteCount:])
	return result
}

// Slice 返回切片的子切片
func Slice(slice []interface{}, start int, end int) []interface{} {
	return slice[start:end]
}

// Replace 替换切片中的元素
func Replace(slice []interface{}, old interface{}, new interface{}) []interface{} {
	result := make([]interface{}, len(slice))
	for i, item := range slice {
		if item == old {
			result[i] = new
		} else {
			result[i] = item
		}
	}
	return result
}

// SortNumber 对切片进行排序
func SortNumber(slice []interface{}, compare func(interface{}, interface{}) bool) []interface{} {
	result := make([]interface{}, len(slice))
	copy(result, slice)
	
	sort.Slice(result, func(i, j int) bool {
		return compare(result[i], result[j])
	})
	
	return result
}

// SortStruct 对结构体切片进行排序
func SortStruct(slice []interface{}, compare func(interface{}, interface{}) bool) []interface{} {
	result := make([]interface{}, len(slice))
	copy(result, slice)
	
	sort.Slice(result, func(i, j int) bool {
		return compare(result[i], result[j])
	})
	
	return result
}

// SortString 对字符串切片进行排序
func SortString(slice []interface{}, compare func(interface{}, interface{}) bool) []interface{} {
	result := make([]interface{}, len(slice))
	copy(result, slice)
	
	sort.Slice(result, func(i, j int) bool {
		return compare(result[i], result[j])
	})
	
	return result
}
