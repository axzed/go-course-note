package chnanel

import "fmt"

func senderV2(ch chan string, down chan struct{}, senderDown chan struct{}) {
	ch <- "hello"
	ch <- "this"
	ch <- "is"
	ch <- "alice"
	// 发送通话结束
	ch <- "EOF"

	// 同步模式等待recver 处理完成
	fmt.Println("sender wait ...")
	<-down
	fmt.Println("sender down ...")

	// sender 退出
	senderDown <- struct{}{}

	// 处理完成后关闭channel
	close(ch)
}

// recver 循环读取chan里面的数据，直到channel关闭
func recverV2(ch chan string, down chan struct{}) {
	defer func() {
		down <- struct{}{}
	}()

	for v := range ch {
		// 处理通话结束
		if v == "EOF" {
			return
		}
		fmt.Println(v)
	}
}

func BufferedChan() {
	ch := make(chan string, 5)

	senderdown := make(chan struct{})
	recverdown := make(chan struct{})
	go senderV2(ch, recverdown, senderdown) // sender goroutine
	go recverV2(ch, recverdown)             // recver goroutine

	// 等待sender执行完成
	<-senderdown
}
