package main

import "fmt"

type T struct {
	s1 string
	m1 map[string]string
}

func (t *T) GetV(k string) string {
	return t.m1[k]
}

func main() {
	t := &T{}
	t.GetV("k1")
	fmt.Println(t.s1)

	// 函数
	Get1(*t)
	fmt.Println(t)

	// 函数做为参数 (高阶函数)
	op := []int{1, 2, 3, 4, 5}
	target := AddThree(op, func(x int) int {
		return x + 3
	})
	fmt.Println(target)

	// 应用   delta(x, y)
	m1 := make(map[string]func(x, y int) int)
	m1["+"] = func(x, y int) int { return x + y }
	m1["-"] = func(x, y int) int { return x - y }
	m1["*"] = func(x, y int) int { return x * y }
	m1["delta"] = func(x, y int) int { return x * y }
}

// t(*T) ---> arg1
// t(T)  ---> arg1
func Get1(arg1 T) {
	arg1.s1 = "m"
}

type Functor func(x int) int

func AddThree(operand []int, fn Functor) (target []int) {
	for _, v := range operand {
		// v += 3
		// v *= 3
		// ...
		target = append(target, fn(v))
	}
	return
}

func Print(e int)

func Add(x, y int) func(e int) {
	return func(input int) {

	}
}
