package day2

import (
	"fmt"
	"testing"
)

func TestIfOne(t *testing.T) {
	/* 局部变量定义 */
	var a int = 100

	/* 判断布尔表达式 */
	if a < 20 {
		/* 如果条件为 true 则执行以下语句 */
		fmt.Printf("a 小于 20\n")
	} else {
		/* 如果条件为 false 则执行以下语句 */
		fmt.Printf("a 不小于 20\n")
	}
	fmt.Printf("a 的值为 : %d\n", a)
}

func TestIfM(t *testing.T) {
	var age int = 18
	if age < 18 {
		fmt.Println("nice")
	} else if age < 28 {
		fmt.Println("beauty")
	} else if age < 38 {
		fmt.Println("sexy")
	} else {
		fmt.Println("next")
	}
}

func TestNested(t *testing.T) {
	var a int = 100
	var b int = 200

	/* 判断条件 */
	if a == 100 {
		/* if 条件语句为 true 执行 */
		if b == 200 {
			/* if 条件语句为 true 执行 */
			fmt.Printf("a 的值为 100 ， b 的值为 200\n")
		}
	}
	fmt.Printf("a 值为 : %d\n", a)
	fmt.Printf("b 值为 : %d\n", b)
}

func TestDeepNested(t *testing.T) {
	// 身高1.8m以上, 25 ~ 35岁, 男
	var (
		height float32
		age    uint
		gender string
		passed bool
	)

	height = 1.9
	age = 30
	gender = "male"

	if height > 1.8 {
		if age > 25 && age <= 35 {
			if gender == "male" {
				passed = true
			}
		}
	}

	if passed {
		fmt.Println("congratulations! your successed")
	} else {
		fmt.Println("not passed")
	}
}

func TestMore1(t *testing.T) {
	// 男, 18~25, 120 ~ 180
	var (
		gender string
		age    uint
		weight uint
	)

	gender = "female"
	age = 20
	weight = 200

	if gender != "male" {
		fmt.Println("male only")
		return
	}

	if age <= 18 {
		fmt.Println("too young")
		return
	}

	if age >= 25 {
		fmt.Println("too old")
		return
	}

	if weight <= 120 || weight >= 180 {
		fmt.Println("too weight or too light")
		return
	}

	fmt.Println("passed")
}

func TestIf2Switch(t *testing.T) {
	gener := "male"

	if gener == "male" {
		fmt.Println("男")
	} else if gener == "female" {
		fmt.Println("女")
	} else {
		fmt.Println("保密")
	}

	switch gener {
	case "male":
		fmt.Println("男")
	case "female":
		fmt.Println("女")
	default:
		fmt.Println("保密")
	}
}

func TestSwich3(t *testing.T) {
	grade := "A"
	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
}

func TestSwichThrouht(t *testing.T) {
	a := 10
	switch a {
	case 10:
		fmt.Println("10")
		fallthrough
	case 20:
		fmt.Println("20")
	default:
		fmt.Println("default")
	}
}

// 1024B = 1KB 1024M= 1G  1024*1024*1024 1M
