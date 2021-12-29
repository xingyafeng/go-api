package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan int, 10)
	done := make(chan bool)

	defer close(message)

	// 1. 消费
	go func() {
		timer := time.NewTicker(time.Second * 1)
		for range timer.C {
			select {
			case ms := <-message:
				fmt.Printf("receiver message: %d\n", ms)
			case <-done: //停止子协程
				fmt.Println("child process interrupt ...")
				return
			default:
				fmt.Printf("#default: receiver message: %d\n", <-message)
			}
		}
	}()

	// 2. 生产
	for i := 0; i < 10; i++ {
		message <- i
	}

	time.Sleep(time.Second * 5)
	close(done) //停止子协程
	time.Sleep(time.Second * 3)
	fmt.Println("main process exit ...")
}
