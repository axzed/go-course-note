package day2

import (
	"fmt"
	"testing"
)

func TestForBase(t *testing.T) {
	var sum int
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func TestForShort(t *testing.T) {
	var sum int
	for sum < 10 {
		sum++
	}

	fmt.Println(sum)
}

func TestForLoop(t *testing.T) {
	var sum int
	for {
		sum++
		fmt.Println(sum)
		if sum == 100 {
			return
		}
	}
}

func TestForRange(t *testing.T) {
	iter := "abcdefg"
	for index, value := range iter {
		fmt.Println(index, value)
		value = 'x'
	}
	fmt.Println(iter)
}

func TestForRangeEdit(t *testing.T) {
	iter := []int{1, 2, 3, 4, 5, 6}
	for index, value := range iter {
		fmt.Println(index, value)
		iter[index] = 99
	}
	fmt.Println(iter)
}

func TestFor99(t *testing.T) {
	for m := 1; m < 10; m++ {
		for n := 1; n <= m; n++ {
			fmt.Printf("%d x %d = %d ", n, m, m*n)
		}
		fmt.Println()
	}
}

func TestForP(t *testing.T) {
	// 遍历 2 ~ 100
	for m := 2; m <= 100; m++ {
		// 假定m是素数 9
		isP := true

		// 判断能不能分解
		for n := 2; n <= (m / n); n++ {
			if m%n == 0 {
				isP = false
				break
			}
		}

		// 无法分解 为素数
		if isP {
			fmt.Println(m)
		}
	}
}

func TestBreak(t *testing.T) {
	i := 0
	for {
		fmt.Println(i)
		if i == 10 {
			break
		}
		i++
	}
}

func TestBreakLalel(t *testing.T) {
	// 不使用标记
	fmt.Println("---- break ----")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break
		}
	}

	// 使用标记
	fmt.Println("---- break label ----")
re:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break re
		}
	}
}

func TestContinue(t *testing.T) {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}

func TestLabelNext(t *testing.T) {
	i := 0
	for {
		if i > 5 {
			goto LABEL
		}
		i++
	}

	fmt.Println("balabala!")

LABEL:
	return
}

func TestLabelPre(t *testing.T) {
	/* 定义局部变量 */
	var a int = 10

	/* 循环 */
LOOP:
	for a < 20 {
		if a == 15 {
			/* 跳过迭代 */
			a++
			goto LOOP
		}
		fmt.Printf("a的值为 : %d\n", a)
		a++
	}
}

func TestForRange1(t *testing.T) {
	a := "abcdefg"
	for i, v := range a {
		fmt.Println(i, v) // rune char
		if v == 'a' {
			v = 'x'
			fmt.Println("change a")
		}
	}
	fmt.Println(a)
}

func TestForBreakContinue(t *testing.T) {
	for i := 1; i <= 10; i++ {
		for k := 1; k <= 10; k++ {
			if k == 5 {
				continue
			}
			fmt.Print(k)
		}
		fmt.Println()
		// fmt.Println(i)
	}
}

func TestFOrLoop1(t *testing.T) {
	i := 1
LOOP:
	for i <= 100 {
		if i == 50 {
			i++
			goto LOOP
		}
		i++
		fmt.Println(i)
	}
}
