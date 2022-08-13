package day1

import (
	"fmt"
	"testing"
)

func TestLocalScope1(t *testing.T) {
	{
		var a string
		a = "local var"
		fmt.Println(a)
	}
	// fmt.Println(a)
}

func TestLocalScope2(t *testing.T) {
	var a string
	a = "parent scope"
	{
		fmt.Println(a)
	}
}

func TestScopeWrite(t *testing.T) {
	var a string
	a = "parent scope"
	{
		var a int
		a = 12
		fmt.Println(a)
	}
	fmt.Println(a)
}

var (
	a = "global parent scope var"
)

func TestGlobal1(t *testing.T) {
	a := 10
	b := 20
	fmt.Println(&a, &b)
	a = b
	fmt.Print(a, b)

}
