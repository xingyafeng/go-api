package main

import (
	"fmt"
	"time"
)

func customer(m chan int, down chan bool) {

	for ch := range m {
		fmt.Println("receiver:", ch)
		fmt.Println("-----")
	}

	fmt.Println("-----")
	// 一个协程只能处理一个通道。
	for range time.NewTicker(time.Second).C {
		select {
		case ch := <-m:
			fmt.Println("receiver:", ch)
			fmt.Println("-----")
		case <-down:
			fmt.Println("child process interrupt ...")
			return
		default:
			fmt.Println("child process default ...")
		}
	}
}

func main() {

	message := make(chan int, 10)
	down := make(chan bool)
	defer close(message)

	// 消费者
	go customer(message, down)

	// 生产者

	for i := 0; i < cap(message); i++ {
		message <- i
		fmt.Println("sender: ", i)
		time.Sleep(time.Second * 2)
	}

	close(down)
	time.Sleep(time.Second * 3)
	fmt.Println("main process exit...")
}
