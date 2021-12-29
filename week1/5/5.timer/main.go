package main

import (
	"fmt"
	"time"
)

/*

	定时器timer

	* timer.Ticker 以指定的时间间隔重复的向通道C发送时间值
	* 使用场景
		* 为协程设定超时时间

 	select 语句的语法：

		每个 case 都必须是一个通信
		所有 channel 表达式都会被求值
		所有被发送的表达式都会被求值
		如果任意某个通信可以进行，它就执行，其他被忽略。
		如果有多个 case 都可以运行，Select 会随机公平地选出一个执行。其他不会执行。
		否则：
			如果有 default 子句，则执行该语句。
			如果没有 default 子句，select 将阻塞，直到某个通信可以运行；
			Go 不会重新对 channel 或值进行求值。

*/

func Channs(ch chan int, stopCh chan bool) {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(time.Second)
		fmt.Println("send")
	}

	stopCh <- true

	defer close(ch)
	defer close(stopCh)
}

func test_timer() {

	// 创建定时器，主要记得要程序结束前停止定时器
	timer := time.NewTicker(time.Second * 1)
	defer timer.Stop()

	ch := make(chan int)
	stopCh := make(chan bool)
	c := 0

	go Channs(ch, stopCh)

	for {
		select {
		case c = <-ch:
			fmt.Println("receiver c: ", c)
		case s := <-ch:
			fmt.Println("receiver s: ", s)
		case t := <-timer.C:
			fmt.Println("timer over: ", t)
		case <-stopCh:
			fmt.Println("end..")
			goto end
		}
	}
end:
}

func main() {
	// test_time_base()
	// test_time_range()
	test_time_range2()

	// test_timer()
}

// 测试基础的time
func test_time_base() {

	// 1 . time.NewTicker
	ticker := time.NewTicker(time.Second * 1)
	for i := 0; i < 3; i++ {
		fmt.Println("time ", i)
		// 延时
		<-ticker.C
	}

	// 重新设置延时
	ticker.Reset(time.Second * 3)

	for i := 0; i < 3; i++ {
		fmt.Println("time ", i)
		// 延时
		<-ticker.C
	}

	fmt.Printf("ticker: %v\n", ticker)

	// 2. time.Timer
	timer := time.NewTimer(time.Second * 5)
	fmt.Printf("timer:  %v\n", timer)

}

// 1 测试 time Ticker
func test_time_range1() {

	t := time.NewTimer(1 * time.Second)
	defer t.Stop()

	go func() {
		for {
			<-t.C
			fmt.Println("default ...")

			// 与2 不同需要重置
			t.Reset(time.Second * 1)
		}
	}()

	time.Sleep(time.Second * 4)
}

//2 测试 time Ticker
func test_time_range2() {

	jobs := []int{10, 3, 6, 2, 2}

	fmt.Println("Ticker Created.")
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	index := 0

	for t := range ticker.C {
		interval := jobs[index]
		fmt.Printf("第 %v 个任务开始执行，收到的时间戳 %v, 时长 %v s\n", index, t, interval)

		// 重置时间，不重置就是定义时间
		// time.Sleep(time.Duration(interval) * time.Second)

		//设置边界保证存在能正常退出循环
		index = index + 1
		// index += 1
		fmt.Println("index:", index)
		if index == len(jobs) {
			break
		}
	}
}
