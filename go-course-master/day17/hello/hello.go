package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"gitee.com/infraboard/go-course/day17/middleware"
)

func hello(wr http.ResponseWriter, r *http.Request) {
	wr.Write([]byte("hello"))
}

func timeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()

		// next handler
		next.ServeHTTP(wr, r)

		timeElapsed := time.Since(timeStart)
		log.Println(timeElapsed)
	})
}

// func main() {
// 	// http.HandleFunc("/", hello)
// 	// HandlerFunc 是一个类型, 我们把hello这个函数，转换为了HandlerFunc类型, 这我们使用int(a)是一个语法 为啥要这样喃?
// 	// 因为 HandlerFunc 实现了ServeHTTP方法, 这样我们的hello函数对象也就有了ServeHTTP方法
// 	// 是不是很骚, HandlerFunc 是个函数， 我们把他定义为一个Type, 然后给这个Type 绑定了一个函数
// 	// 这样凡事 被转换为HandlerFunc的类型，都是一个http.Handler
// 	// 当然 要完成 Type()的转换，函数签名必须一致
// 	http.Handle("/", timeMiddleware(http.HandlerFunc(hello)))
// 	err := http.ListenAndServe(":8080", nil)
// 	fmt.Println(err)
// }

func main() {
	r := middleware.NewRouter()
	r.Use(middleware.LogMiddleware)
	r.Use(middleware.TimeMiddleware)
	http.Handle("/", r.Merge(http.HandlerFunc(hello)))
	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
}
