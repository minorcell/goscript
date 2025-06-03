// Package goscript 提供了类似 JavaScript 的函数操作方法
// 所有的实现都在 packages 子目录中，但通过此文件导出到顶层包
package goscript

import "github.com/minorcell/goscript/packages"

// 切片操作函数
// ----------------------------

// Filter 过滤一个切片，返回符合条件的元素
func Filter(slice []interface{}, predicate func(interface{}) bool) []interface{} {
	return packages.Filter(slice, predicate)
}

// Map 对切片中的每个元素应用一个转换函数
func Map(slice []interface{}, mapper func(interface{}) interface{}) []interface{} {
	return packages.Map(slice, mapper)
}

// Reduce 将切片中的元素归约为单个值
func Reduce(slice []interface{}, initialValue interface{}, reducer func(interface{}, interface{}) interface{}) interface{} {
	return packages.Reduce(slice, initialValue, reducer)
}

// 字符串操作函数
// ----------------------------

// Contains 检查一个字符串是否包含子字符串
func Contains(s, substr string) bool {
	return packages.Contains(s, substr)
}

// Split 将字符串按分隔符分割成切片
func Split(s, sep string) []string {
	return packages.Split(s, sep)
}

// Join 将字符串切片用分隔符连接成一个字符串
func Join(elems []string, sep string) string {
	return packages.Join(elems, sep)
}

// Trim 去除字符串两端的空白字符
func Trim(s string) string {
	return packages.Trim(s)
}
