package main

import (
	"fmt"
	"time"
)

func main() {
	decorator(Hello)("boy")
}

func Hello(s string) {
	fmt.Printf("hello, %s\n", s)
}

func decorator(fn func(s string)) func(s string) {
	return func(s string) {
		start := time.Now()
		fn(s)
		delta := time.Now().Sub(start)
		fmt.Println(delta)
	}
}
