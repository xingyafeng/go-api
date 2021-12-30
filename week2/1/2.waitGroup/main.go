package main

import (
	"fmt"
	"sync"
	"time"
)

// 线程安全，鼓励使用channel
// Sync
//
// 保证 前面的程序全部执行完成，再执行后面的程序，当程序是一个耗时该如何保证？
func main() {

	waitByWG()
}

func waitBySleep() {

	for i := 0; i < 100; i++ {
		println(i)
	}

	time.Sleep(time.Duration(1000 * time.Microsecond))
}

// channel 的使用
// 注意:
// 1. 不能阻塞主线程，当主线程发生阻塞的时候，会出现如下错误:
//      all goroutines are asleep - deadlock-Go中协程死锁
// 2. 缓冲问题要特别注意哈
// 	  当写10 读小于10 不会出现死锁，当读大于10就出现问题啦。
func waitByChannel() {

	ch := make(chan int)
	for i := 0; i < 31; i++ {
		go func(i int) {
			fmt.Println("i:", i)
			ch <- i
		}(i)
	}

	for i := 0; i < 1; i++ {
		// go func(i int) {
		ret := <-ch
		println(ret)
		// }(i)
	}

	time.Sleep(time.Second)
}

// 最正确的做法。
// sync.WaitGroup 含义， 通过Add 方法来保证程序循环的次数，当次数未到时一直阻塞。
// 直到执行循环完成才能继续执行下面的程序
// 实现原来：count 一直减到 0
func waitByWG() {

	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int, wg *sync.WaitGroup) {
			fmt.Println(i)
			wg.Done()
		}(i, &wg)
	}
	wg.Wait()
}
