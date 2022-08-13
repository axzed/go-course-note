package day3

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestUnsafeMain(t *testing.T) {
	a := [3]int64{1, 2, 3}
	fmt.Printf("%p\n", &a)

	s1 := unsafe.Sizeof(a[0])
	fmt.Printf("%d\n", s1) // a + 8B  a[1]

	// 1. *[3]int64 --> unsafe.Pointer ---> uintptr(uint64)
	p1 := &a // *p1
	fmt.Println(*(*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(p1)) + unsafe.Sizeof(a[0]))))
	// fmt.Println() 标准输入
	// int int32 float32  *int *int32 *float64    -->   T -->
	var x int64 = 10
	y := int8(x)                   // int64 ---> int8
	x = int64(y)                   // int8  ---> int64
	x1 := &x                       // x1 *int64
	unsafeP1 := unsafe.Pointer(x1) // *int64 ---> Pointer
	fmt.Println(unsafeP1)
	intPtr1 := (*int64)(unsafeP1) // Pointer ---> *int64
	fmt.Println(intPtr1)
}

func TestSafePointer1(t *testing.T) {
	str := "pointer_test"
	a := &str
	fmt.Println(a)
}

func TestPonterT(t *testing.T) {
	var x int64 = 20
	a := &x // *int64
	fmt.Printf("%T", a)
	// 没有办法指针a的值进行如下转换: *int64 --> *uint64
}

func TestPonterT1(t *testing.T) {
	var x int64 = 20
	a := &x
	fmt.Println(a)

	var y uint64 = 20
	b := &y
	fmt.Println(b)

	// 我们不能进行 a = b
}

type Man struct {
	Name string
	Age  int64
}

func TestUnsafePonter1(t *testing.T) {
	m := Man{Name: "John", Age: 20}
	fmt.Printf("%p\n", &m)
	fmt.Println(unsafe.Sizeof(m.Name), unsafe.Sizeof(m.Age), unsafe.Sizeof(m)) // 16 8 24
	fmt.Println(unsafe.Offsetof(m.Name))                                       // 0
	fmt.Println(unsafe.Offsetof(m.Age))
}

func TestUnsafePointer2(t *testing.T) {
	a := [3]int64{1, 2, 3}
	fmt.Printf("%p\n", &a)

	s1 := unsafe.Sizeof(a[0])
	fmt.Printf("%d\n", s1)

	p1 := (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(&a)) + s1))
	fmt.Println(*p1)
}

func TestUnsafePointer3(t *testing.T) {
	type T struct{ a int }
	var t1 T
	fmt.Printf("%p\n", &t1)                          // 0xc0000a0200
	println(&t1)                                     // 0xc0000a0200
	fmt.Printf("%x\n", uintptr(unsafe.Pointer(&t1))) // c0000a0200
}

type T struct {
	x bool
	y [3]int16 // 1 index
}

const (
	N = unsafe.Offsetof(T{}.y)
	M = unsafe.Sizeof(T{}.y[0])
)

func TestUnsafePointer4(t *testing.T) {
	t1 := T{y: [3]int16{123, 456, 789}}
	p := unsafe.Pointer(&t1)
	// "uintptr(p) + N + M + M"为t.y[2]的内存地址。 offset
	ty2 := (*int16)(unsafe.Pointer(uintptr(p) + N + M))
	fmt.Println(*ty2) // 789
}
