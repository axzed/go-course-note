// package doc

package main

import (
	"fmt"
)

var (
	hello = "欢迎光临, xxxxx"
)

func main() {
	a := "hello"           // var a string; a = "hello"  "1" 10, 100.1
	fmt.Printf("%p\n", &a) // &运算符  8Byte uint64 内存地址
	fmt.Println(a)
	fmt.Println([]byte(a)) // T(a) --> [104 101 108 108 111]
	// a             --->      hello
	// 0xc00004e240             值
	// 8Byte uint64            5Byte
	// 获取变量的值:  内存地址的访问
	// 0xc00004e240 -->  5Byte(hello)

	// a := "hello", a= 10  a = 10.1
	fmt.Println(hello)
	fmt.Println(hello)
	fmt.Println(hello) // {} 词法块
	{
		{
			b := "hello"
			fmt.Printf("%p\n", &b)
			{
				b := "private b"
				fmt.Println(b)
				fmt.Printf("%p\n", &b)

			}
			fmt.Println(b)
		}
	}
}
