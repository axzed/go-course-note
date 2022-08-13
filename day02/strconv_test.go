package day2

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"unsafe"
)

func TestTypeOf1(t *testing.T) {
	a := 10
	b := 0.1314
	c := "hello"

	fmt.Printf("a type: %v\n", reflect.TypeOf(a))
	fmt.Printf("b type: %v\n", reflect.TypeOf(b))
	fmt.Printf("c type: %v\n", reflect.TypeOf(c))
}

func TestF2Int(t *testing.T) {
	a := 3.14
	fmt.Println(int(a))
}

func TestTypeOf(t *testing.T) {
	type Age int
	var a Age = 10
	var b int = 20
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
}

func TestCustom(t *testing.T) {
	// Age 底层数据结构为 int
	type Age int
	// a 类型是Age 底层为 int 10
	var a Age = 10

	// 将a转化成int类型,
	// 由于a是Age, 转化成int后, 他们不是同一种类型，不能再次赋值回去: a = int(a) 是不行的
	b := int(a)
	// 现在b是int类型
	fmt.Println(reflect.TypeOf(b))

	// 反过来我们也可以将int类型转换为Age类型
	c := Age(10)
	// 现在c就是Age类型，而不是int类型了
	fmt.Println(reflect.TypeOf(c))
}

func TestConvD(t *testing.T) {
	var a float64 = 5.3232223232323
	fmt.Println(float32(a))
	fmt.Println(int(a))
	fmt.Println(int8(a))
	fmt.Println(int16(a))
	fmt.Println(int32(a))
	fmt.Println(int64(a))
	fmt.Println(uint(a))
}

func TestItoa(t *testing.T) {
	str := strconv.Itoa(100)
	fmt.Printf("type %v, value: %s\n", reflect.TypeOf(str), str)
}

func TestPAtoi(t *testing.T) {
	i, err := strconv.Atoi("100")
	fmt.Printf("type %v, value: %d, err: %v\n-", reflect.TypeOf(i), i, err)
}

func TestPAtoiErr(t *testing.T) {
	i, err := strconv.Atoi("100x")
	fmt.Printf("type %v, value: %d, err: %v\n-", reflect.TypeOf(i), i, err)
}

func TestParseBool(t *testing.T) {
	b, err := strconv.ParseBool("true")
	fmt.Println(b, err)
}

func TestParseFloat(t *testing.T) {
	f1, err := strconv.ParseFloat("3.1", 32)
	fmt.Println(f1, err)
	f2, err := strconv.ParseFloat("3.1", 64)
	fmt.Println(f2, err)
}

func TestFloat1(t *testing.T) {
	var n float64 = 0
	for i := 0; i < 1000; i++ {
		n += .01
	}
	fmt.Println(n)
}

func bInt8(n int8) string {
	fmt.Println(*(*uint8)(unsafe.Pointer(&n))) // 1111 1111
	return strconv.FormatUint(uint64(*(*uint8)(unsafe.Pointer(&n))), 2)
}

func TestParseInt(t *testing.T) {
	fmt.Println(bInt8(-1)) // 0000 0001(原码) -> 1111 1110(反码) -> 1111 1111(补码)
	// Parse 二进制字符串
	i, err := strconv.ParseInt("11111111", 2, 16)
	fmt.Println(i, err)
	// Parse 十进制字符串
	i, err = strconv.ParseInt("-255", 10, 16)
	fmt.Println(i, err)
	// Parse 十六进制字符串
	i, err = strconv.ParseInt("4E2D", 16, 16)
	fmt.Println(i, err)
}

func TestParseUint(t *testing.T) {
	u, err := strconv.ParseUint("11111111", 2, 16)
	fmt.Println(u, err)
	u, err = strconv.ParseUint("255", 10, 16)
	fmt.Println(u, err)
	u, err = strconv.ParseUint("4E2D", 16, 16)
	fmt.Println(u, err)
}

func TestFormatInt(t *testing.T) {
	fmt.Println(strconv.FormatBool(true))

	// 问题又来了
	fmt.Println(strconv.FormatInt(255, 2))
	fmt.Println(strconv.FormatInt(255, 10))
	fmt.Println(strconv.FormatInt(255, 16))

	fmt.Println(strconv.FormatUint(255, 2))
	fmt.Println(strconv.FormatUint(255, 10))
	fmt.Println(strconv.FormatUint(255, 16))

	fmt.Println(strconv.FormatFloat(3.1415, 'E', -1, 64))
}

func TestStrconLost(t *testing.T) {
	var n float64 = 0
	for i := 0; i < 100; i++ {
		n += .01
		fmt.Println(n)
	}

	fmt.Println(float64(0.05) + float64(0.01))
}
