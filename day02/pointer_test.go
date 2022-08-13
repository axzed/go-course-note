package day2

import (
	"fmt"
	"testing"
)

func TestPointerRaw(t *testing.T) {
	a := "string a"
	fmt.Println(&a)
}

func TestPointer1(t *testing.T) {
	var a *int
	fmt.Println(a)
}

func TestPointer2(t *testing.T) {
	var a *int = new(int)
	fmt.Println(a)
	fmt.Println(*a)
}

func TestPointer3(t *testing.T) {
	var a *int = new(int)
	*a = 10
	fmt.Println(a, *a)
}

func TestPPP(t *testing.T) {
	var a ****int

	v := 10
	p1 := &v
	p2 := &p1
	p3 := &p2
	a = &p3
	fmt.Println(v, p1, p2, p3, a)
}

func TestPointer11(t *testing.T) {
	var ip *int = new(int)
	fmt.Println(&ip, ip, *ip)
}
