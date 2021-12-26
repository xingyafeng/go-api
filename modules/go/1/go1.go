package main

import (
	"fmt"
	"time"
)

func test_time() {

	// 1 . time.NewTicker
	ticker := time.NewTicker(time.Second)
	<-ticker.C
	ticker.Reset(time.Second * 5)
	ticker.Stop()
	fmt.Printf("ticker: %v\n", ticker)

	// 2. time.Timer
	timer := time.NewTimer(time.Second * 5)
	fmt.Printf("timer:  %v\n", timer)

}

func test_xx() {

	jobs := []int{10, 3, 6, 2, 2}

	fmt.Println("Ticker Created.")
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	index := 0

	for t := range ticker.C {
		interval := jobs[index]
		fmt.Printf("第 %v 个任务开始执行，收到的时间戳 %v, 时长 %v s\n", index, t, interval)
		time.Sleep(time.Duration(interval) * time.Second)
		index = index + 1
		if index == len(jobs) {
			break
		}
	}
}

func main() {
	// test_time()
	test_xx()
	return
	ch := make(chan int, 10)

	// producer
	ticker := time.NewTicker(time.Second)

	go func(ticker *time.Ticker) {
		for i := 0; i < 3; i++ {
			<-ticker.C
			ch <- i
		}
		defer ticker.Stop()
		defer close(ch)
	}(ticker)

	// consumer
	for item := range ch {
		fmt.Printf("recived %d\n", item)
	}
}
