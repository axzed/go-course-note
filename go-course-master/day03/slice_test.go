package day3

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestSliceMain(t *testing.T) {
	s1 := make([]int, 0, 4)
	s1 = append(s1, 10, 20, 30, 40) // 10, 20, 30, 40
	fmt.Println(s1, len(s1), cap(s1))

	fmt.Println(sum1(s1))
}

func sum1(args []int) int {
	sum := 0
	for _, v := range args {
		sum += v
	}
	return sum
}

func TestSlice1(t *testing.T) {
	slice := []int{3: 100}
	fmt.Println(slice, len(slice), cap(slice))
}

func TestSliceSize(t *testing.T) {
	a := make([]int, 3, 5)
	fmt.Println(a, len(a), cap(a))
	fmt.Println(a[2])
	fmt.Println(a, len(a), cap(a))
}

func TestSliceAddr(t *testing.T) {
	a := make([]int, 3, 5)
	fmt.Printf("%p\t", a)
	fmt.Println(&(a[0]), &(a[1]), 0xa3c0-0xa3c8)

	b := make([]int64, 3, 5)
	fmt.Println(unsafe.Sizeof(&(b[0])))
	fmt.Println(&(b[0]), &(b[1]), 0xa3c0-0xa3c8)
}

func TestSlice3(t *testing.T) {
	var s []int
	s[0] = 1
}

func TestSlice4(t *testing.T) {
	s := []int{1, 2, 3}
	_ = s[3]
}

func TestSliceAppend1(t *testing.T) {
	s := make([]int, 0, 4)
	s = append(s, 10, 20, 30, 40)
	fmt.Println(s)
}

func TestSlice5(t *testing.T) {
	s1 := []int{1, 2, 3, 4}
	s2 := s1[2:4:4]
	fmt.Println(s2, len(s2), cap(s2))

	s3 := s1[2:4]
	fmt.Println(s3, len(s3), cap(s3))
}

func TestSlice6(t *testing.T) {
	s1 := []int{10, 20, 30, 40}
	s2 := s1[1:3]
	fmt.Println(s2, len(s2), cap(s2))

	fmt.Println(s1[1], s2[0])
	s1[1] = 200
	fmt.Println(s1[1], s2[0])
}

func TestSlice7(t *testing.T) {
	s1 := []int{10, 20, 30, 40}
	s2 := s1[1:3:3]
	fmt.Println(s2, len(s2), cap(s2))

	fmt.Println(s1[1], s2[0])
	s2 = append(s2, 30)
	s1[1] = 200
	fmt.Println(s1[1], s2[0])
}

func TestCopy1(t *testing.T) {
	s1 := []int{10, 20, 30, 40}
	s2 := make([]int, 5)
	num := copy(s2, s1)
	fmt.Println(num)
	fmt.Println(s1, s2)
}

func TestSliceNil(t *testing.T) {
	s := make([]int, 0)
	s[0] = 1
}
