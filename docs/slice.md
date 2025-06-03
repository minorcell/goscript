# Slice 使用说明

> Golang 中的切片，主要提供了类似于 JavaScript 中的数组操作方法。

## Filter

> Filter 过滤一个切片，返回符合条件的元素，不对原切片进行修改

```go
filter(slice []interface{}, predicate func(interface{}) bool) []interface{}
```

示例：

```go
func main() {
	slice := []interface{}{1, 2, 3, 4, 5}
	filtered := Filter(slice, func(x interface{}) bool {
		return x.(int)%2 == 0
	})
	fmt.Println(filtered) // [2 4]
}
```

## Map

> Map 对切片中的每个元素应用一个转换函数，返回转换后的切片，不对原切片进行修改

```go
map(slice []interface{}, mapper func(interface{}) interface{}) []interface{}
```

示例：

```go
func main() {
	slice := []interface{}{1, 2, 3, 4, 5}
	mapped := Map(slice, func(x interface{}) interface{} {
		return x.(int) * 2
	})
	fmt.Println(mapped) // [2 4 6 8 10]
}
```

## Reduce

> Reduce 将切片中的元素归约为单个值，返回单个值，不对原切片进行修改

```go
reduce(slice []interface{}, initialValue interface{}, reducer func(interface{}, interface{}) interface{}) interface{}
```

示例：

```go
func main() {
	slice := []interface{}{1, 2, 3, 4, 5}
	reduced := Reduce(slice, 0, func(x, y interface{}) interface{} {
		return x.(int) + y.(int)
	})
	fmt.Println(reduced) // 15
}
```
