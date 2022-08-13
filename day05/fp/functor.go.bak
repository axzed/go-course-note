package main

import "fmt"

func main() {
	op := []int{1, 2, 3, 4, 5}
	target := AddThree(op, func(x int) int {
		return x + 3
	})
	fmt.Println(target)
}

type Functor func(x int) int

func AddThree(operand []int, fn Functor) (target []int) {
	for _, v := range operand {
		target = append(target, fn(v))
	}
	return
}
