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

## Max

> Max 返回切片中的最大值，不对原切片进行修改

```go
max(slice []interface{}) interface{}
```

示例：

```go
func main() {
	slice := []interface{}{1, 2, 3, 4, 5}
	max := Max(slice)
	fmt.Println(max) // 5
}
```

## Min

> Min 返回切片中的最小值，不对原切片进行修改

```go
min(slice []interface{}) interface{}
```

示例：

```go
func main() {
	slice := []interface{}{1, 2, 3, 4, 5}
	min := Min(slice)
	fmt.Println(min) // 1
}
```

## Splice

> Splice 替换切片中的元素，返回替换后的切片，不对原切片进行修改

```go
splice(slice []interface{}, start int, deleteCount int, items ...interface{}) []interface{}
```

示例：

```go
func main() {
	slice := []interface{}{1, 2, 3, 4, 5}
	spliced := Splice(slice, 1, 2, 6, 7)
	fmt.Println(spliced) // [1 6 7 4 5]
}
```

## Slice

> Slice 返回切片的子切片，不对原切片进行修改

```go
slice(slice []interface{}, start int, end int) []interface{}
```

示例：

```go
func main() {
	slice := []interface{}{1, 2, 3, 4, 5}
	subSlice := Slice(slice, 1, 3)
	fmt.Println(subSlice) // [2 3]
}
```

## Replace

> Replace 替换切片中的元素，返回替换后的切片，不对原切片进行修改

```go
replace(slice []interface{}, old interface{}, new interface{}) []interface{}
```

示例：

```go
func main() {
	slice := []interface{}{1, 2, 3, 4, 5}
	replaced := Replace(slice, 2, 6)
	fmt.Println(replaced) // [1 6 3 4 5]
}
```

## SortNumber

> SortNumber 对切片进行排序，返回排序后的切片，不对原切片进行修改

```go
sortNumber(slice []interface{}, compare func(interface{}, interface{}) bool) []interface{}
```

示例：

```go
func main() {
	slice := []interface{}{5, 2, 3, 1, 4}
	sorted := SortNumber(slice, func(x, y interface{}) bool {
		return x.(int) < y.(int)
	})
	fmt.Println(sorted) // [1 2 3 4 5]
}
```

```go
func main() {
	slice := []interface{}{5, 2, 3, 1, 4}
	sorted := SortNumber(slice, func(x, y interface{}) bool {
		return x.(int) > y.(int)
	})
	fmt.Println(sorted) // [5 4 3 2 1]
}
```
	
## SortStruct

> SortStruct 对结构体切片进行排序，返回排序后的切片，不对原切片进行修改

```go
sortStruct(slice []interface{}, compare func(interface{}, interface{}) bool) []interface{}
```

示例：

```go
func main() {
	type Person struct {
		Name string
		Age  int
	}
	persons := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}
	sorted := SortStruct(persons, func(x, y interface{}) bool {
		return x.(Person).Age < y.(Person).Age
	})
	fmt.Println(sorted) // [{Bob 25} {Alice 30} {Charlie 35}]
}
```

```go
func main() {
	type Person struct {
		Name string
		Age  int
	}
	persons := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}
	sorted := SortStruct(persons, func(x, y interface{}) bool {
		return x.(Person).Age > y.(Person).Age
	})
	fmt.Println(sorted) // [{Charlie 35} {Alice 30} {Bob 25}]
}
```

## SortString

> SortString 对字符串切片进行排序，返回排序后的切片，不对原切片进行修改

```go
sortString(slice []interface{}, compare func(interface{}, interface{}) bool) []interface{}
```

示例：

```go
func main() {
	slice := []interface{}{"apple", "banana", "cherry"}
	sorted := SortString(slice, func(x, y interface{}) bool {
		return x.(string) < y.(string)
	})
	fmt.Println(sorted) // [apple banana cherry]
}
```

```go
func main() {
	slice := []interface{}{"apple", "banana", "cherry"}
	sorted := SortString(slice, func(x, y interface{}) bool {
		return x.(string) > y.(string)
	})
	fmt.Println(sorted) // [cherry banana apple]
}
```

