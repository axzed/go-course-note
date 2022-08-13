package day3

import (
	"fmt"
	"testing"
)

func TestPanic1(t *testing.T) {
	fn()
}

func fn() {
	fmt.Println("start fn")
	panic("pannic in fn")
	fmt.Println("end fn")
}

func TestPanic2(t *testing.T) {
	var a *int
	fmt.Println(*a)
}

func TestPanic3(t *testing.T) {
	var x, y *int
	sum(x, y)
}

func sum(x, y *int) int {
	return *x + *y
}

func TestPanic4(t *testing.T) {
	var x, y *int
	fmt.Println(recover())
	sum(x, y)
}

func TestPanic5(t *testing.T) {
	defer func() {
		fmt.Println(recover())
	}()

	var x, y *int
	sum(x, y)
}

type MyType func(x, y int) int

func Test11(t *testing.T) {
	var my MyType
	my = func(x, y int) int {
		return x + y
	}

	fmt.Println(T11(my))
}

func T11(fn MyType) int {
	return fn(1, 2) * 2
}

func TestFact(t *testing.T) {
	fmt.Println(fact(4))
}

// // n*(n-1)*...*3*2*1 5 * fact(4)
func fact(n int) int {
	if n == 1 {
		return 1
	}

	return n * fact(n-1) // n * fact(n-1) * fact(n-2) * 1 <退出点> 3* 2* *1
}

func TestFib(t *testing.T) {
	fmt.Println(fib(3))
}

// f(n) = f(n-1) + f(n-2)  // n == 2     1, 1, 2, 3, 5, 8, 13 (13)
// n == 2
func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
