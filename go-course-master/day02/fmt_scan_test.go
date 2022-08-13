package day2

import (
	"fmt"
	"testing"
)

var (
	name string
	age  uint
)

func TestSscan(t *testing.T) {
	input := "andy 30"
	fmt.Sscan(input, &name, &age)
	fmt.Println(name, age)
}

func TestSscan2(t *testing.T) {
	input := "andy 30 40"
	fmt.Sscan(input, &name, &age)
	fmt.Println(name, age)
}

func TestSscan3(t *testing.T) {
	input := "andy\n30"
	fmt.Sscan(input, &name, &age)
	fmt.Println(name, age)
}

func TestSscanln(t *testing.T) {
	input := "andy 30 40"
	fmt.Sscanln(input, &name, &age)
	fmt.Println(name, age)
}

func TestSscanln2(t *testing.T) {
	input := "andy \n 30 40"
	fmt.Sscanln(input, &name, &age)
	fmt.Println(name, age)
}

func TestSscanf(t *testing.T) {
	input := "andy : 50"
	fmt.Sscanf(input, "%s : %d", &name, &age)
	fmt.Println(name, age)
}

func TestSscanRet(t *testing.T) {
	input := "andy : 50"
	ret, err := fmt.Sscanf(input, "%s : %d", &name, &age)
	fmt.Println(ret, err)
}

func TestSscanRet2(t *testing.T) {
	input := "andy : err"
	ret, err := fmt.Sscanf(input, "%s : %d", &name, &age)
	fmt.Println(ret, err)
}
