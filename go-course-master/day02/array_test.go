package day2

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	payload := [4]int{1}
	fmt.Printf("%p\n", &payload)
	change(payload)
}

func change(payload [4]int) {
	fmt.Printf("%p\n", &payload)
	payload[0] = 10
}

func TestMaxint(t *testing.T) {
	var a []int
	a = append(a, 1)
	fmt.Println(a)
}

func TestArrayType(t *testing.T) {
	var (
		a1 [4]int
		a2 [5]int
	)
	fmt.Println(reflect.TypeOf(a1))
	fmt.Println(reflect.TypeOf(a2))
}

func TestArrayD(t *testing.T) {
	var a1 [5]int
	fmt.Println(a1)

	var a2 [4]string
	fmt.Println(a2)
}

func TestArrayD1(t *testing.T) {
	a1 := [3]int{1, 2, 3}
	fmt.Println(a1)
	// 如果将元素个数指定为特殊符号...，则表示通过初始化时的给定的值个数来推断数组长度
	a2 := [...]int{1, 2, 3, 4}
	fmt.Println(a2)
	a3 := [...]int{1, 1, 1}
	fmt.Println(a3)
	// 如果声明数组时，只想给其中某几个元素初始化赋值，则使用索引号
	a4 := [4]int{0: 1, 3: 5}
	fmt.Println(a4)
}

func TestArrayV1(t *testing.T) {
	a := [4]int{0: 1, 3: 5}
	fmt.Println(a[0])
	fmt.Println(a[3])

	a[0] = 10
	a[3] = 20
	fmt.Println(a[0])
	fmt.Println(a[3])
}

func TestArrayPtr(t *testing.T) {
	a := [4]*int{0: new(int), 3: new(int)}
	fmt.Println(a)

	// 如果指针地址为空, 是会报空指针错误的, 比如
	// *a[1] = 3

	*a[0] = 10
	*a[3] = 20
	fmt.Println(a)
	fmt.Println(*a[0], *a[3])

	// 为1赋值
	a[1] = new(int)
	*a[1] = 30
	fmt.Println(a, *a[1])
}

func TestArrayCopy(t *testing.T) {
	a1 := [4]string{"a", "b", "c", "m"}
	a2 := [4]string{"x", "y", "z", "n"}
	a1 = a2
	fmt.Println(a1, a2)
}

func TestArrayPtrCopy(t *testing.T) {
	a1 := [4]*string{new(string), new(string), new(string), new(string)}
	a2 := a1
	fmt.Println(a1, a2)

	*a1[0] = "A"
	*a1[1] = "B"
	*a1[2] = "C"
	*a1[3] = "C"
	fmt.Println(*a1[0], *a2[0])
}

func TestArrayIter(t *testing.T) {
	a := [4]int{1, 2, 3, 4}
	for i, v := range a {
		fmt.Println(i, v)
	}
}

func TestArray2x(t *testing.T) {
	pos := [4][2]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}}
	fmt.Println(pos)

	pos[0] = [2]int{10, 10}
	fmt.Println(pos)
}

func TestArray(t *testing.T) {
	a := [2]int{}
	fmt.Println(a, a[1])
}
