# String 使用说明

> Golang 中的字符串，主要提供了类似于 JavaScript 中的字符串操作方法。

## Contains

> Contains 检查一个字符串是否包含子字符串，返回 bool 值，不对原字符串进行修改

```go
contains(s, substr string) bool
```

示例：

```go
func main() {
	s := "Hello, World!"
	fmt.Println(Contains(s, "World")) // true
	fmt.Println(Contains(s, "Universe")) // false
}
```

## Split

> Split 将字符串按分隔符分割成切片，返回 []string，不对原字符串进行修改

```go
split(s, sep string) []string
```

示例：

```go
func main() {
	s := "Hello, World!"
	split(s, " ") // ["Hello," "World!"]
}
```

## Join

> Join 将字符串切片用分隔符连接成一个字符串，返回 string，不对原切片进行修改

```go
join(elems []string, sep string) string
```

示例：

```go
func main() {
	elems := []string{"Hello", "World!"}
	join(elems, " ") // "Hello World!"
}
```
