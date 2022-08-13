package main

import (
	"bufio"
	"fmt"
	"os"

	"gitee.com/infraboard/go-course/day1"
)

var (
	name string
	age  uint
)

func Scan() {
	day1.Day1T()
	fmt.Print("请输入你的姓名和年龄: ")
	fmt.Scan(&name, &age)
	fmt.Printf("姓名: %s 年龄: %d", name, age)
	fmt.Println()
}

func Scanln() {
	fmt.Print("请输入你的姓名和年龄, 以空格分隔: ")
	fmt.Scanln(&name, &age)
	fmt.Printf("姓名: %s 年龄: %d", name, age)
	fmt.Println()
}

func Scanf() {
	fmt.Print("请输入你的姓名和年龄, 以 ：分隔: ")
	fmt.Scanf("%s : %d", &name, &age)
	fmt.Printf("姓名: %s 年龄: %d", name, age)
	fmt.Println()
}

func ScanfFromBufio() {
	fmt.Print("请输入你的姓名和年龄, 以 : 分隔: ")
	stdin := bufio.NewReader(os.Stdin)
	line, _, err := stdin.ReadLine()
	if err != nil {
		panic(err)
	}
	n, err := fmt.Sscanf(string(line), "%s : %d", &name, &age)
	if err != nil {
		panic(err)
	}
	fmt.Printf("read number of items: %d", n)
	fmt.Println()
	fmt.Printf("姓名: %s 年龄: %d", name, age)
	fmt.Println()
}

func main() {
	ScanfFromBufio()
}
