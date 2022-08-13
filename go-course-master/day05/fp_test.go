package day4

import (
	"fmt"
	"strings"
	"testing"
)

func TestMap(t *testing.T) {
	list := []string{"abc", "def", "fqp"}
	out := MapStrToUpper(list, func(item string) string {
		return strings.ToUpper(item)
	})
	fmt.Println(out)
}

func MapStrToUpper(data []string, fn func(string) string) []string {
	newData := make([]string, 0, len(data))
	for _, v := range data {
		newData = append(newData, fn(v))
	}

	return newData
}

func TestReduce(t *testing.T) {
	list := []string{"abc", "def", "fqp", "abc"}
	// 统计字符数量
	out1 := ReduceSum(list, func(s string) int {
		return len(s)
	})
	fmt.Println(out1)
	// 出现过ab的字符串数量
	out2 := ReduceSum(list, func(s string) int {
		if strings.Contains(s, "ab") {
			return 1
		}
		return 0
	})
	fmt.Println(out2)
}

func ReduceSum(data []string, fn func(string) int) int {
	sum := 0
	for _, v := range data {
		sum += fn(v)
	}
	return sum
}

func TestFilter(t *testing.T) {
	list := []string{"abc", "def", "fqp", "abc"}
	out := ReduceFilter(list, func(s string) bool {
		return strings.Contains(s, "f")
	})
	fmt.Println(out)
}

func ReduceFilter(data []string, fn func(string) bool) []string {
	newData := []string{}
	for _, v := range data {
		if fn(v) {
			newData = append(newData, v)
		}
	}
	return newData
}
