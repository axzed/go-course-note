package main

import "fmt"

func main() {
	var a ****int

	v := 10
	p1 := &v
	p2 := &p1
	p3 := &p2
	a = &p3

	fmt.Println(a, *a, **a, ***a, ****a)
}
