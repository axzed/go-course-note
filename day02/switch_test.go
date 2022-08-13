package day2

import (
	"fmt"
	"testing"
)

const (
	Unknown = iota
	Male
	Female
)

func TestIfEnum(t *testing.T) {
	gender := 0

	if gender == Unknown {

	} else if gender == Male {

	} else if gender == Female {

	} else {
		fmt.Println()
	}
}

func TestIfToSwitch(t *testing.T) {
	gender := 0
	switch gender {
	case Unknown:
	case Male:
	case Female:
	default:
		fmt.Println()
	}
}

func TestSwitch1(t *testing.T) {
	/* 定义局部变量 */
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

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
	fmt.Printf("你的等级是 %s\n", grade)
}

func TestFallthrough(t *testing.T) {
	switch {
	case false:
		fmt.Println("1、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	case false:
		fmt.Println("3、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("4、case 条件语句为 true")
	case false:
		fmt.Println("5、case 条件语句为 false")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}
}

func TestWorkFlow(t *testing.T) {
	step := 0

	switch step {
	case 0:

	case 1:
	case 2:
	default:

	}
}
