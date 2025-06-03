package packages

import "strings"

// Contains 检查一个字符串是否包含子字符串
func Contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// Split 将字符串按分隔符分割成切片
func Split(s, sep string) []string {
	return strings.Split(s, sep)
}

// Join 将字符串切片用分隔符连接成一个字符串
func Join(elems []string, sep string) string {
	return strings.Join(elems, sep)
}

// Trim 去除字符串两端的空白字符
func Trim(s string) string {
	return strings.TrimSpace(s)
}
