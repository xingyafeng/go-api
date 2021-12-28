package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	channel ---- 多线程间通讯

		* 一段发送数据 一段接受数据
		* 同一时间只有一个协程可以访问数据，无共享内存模式可能出现的内存竞争
		* 协调协程的执行顺序

	申明方式
		* var identifier chan datatype
		* 操作符 <-

	用途：
		* 1、多个协程之间传递数据的时候
		* 2、多个协程之前需要协调执行顺序的时候

	通道缓冲区

		* 基于channel的通信是同步的
		* 当缓冲区满时，数据的发送是阻塞的
		* 通过make 关键字创建通道时，可以定义缓存区容量。默认缓冲区的容量是 -0-

	通道取数据的方法：
		1、var ret = <-ch
		2、for v:= range ch {
			fmt.Println("v: ", v)
		}

	单向通道
		* var sendOnly chan <- int # 只发送
		* var readOnly <- chan int # 只接受
		* func (w *WebhookCertPatcher) runWebhookController(stopChan <-struct{}) {}

	双向通道 如何使用呢？

	关闭通道
		* defer close(ch)
		* 关闭通道的作用： 告诉接受者该通道无新数据发送
		* 只有发送方才需要关闭通道

	select
		* 当多个协程同时运行时，可通过select轮询多个通道
			* 如果所有通道都阻塞等待，如定义了default,则执行default
			* 如多个通道就绪就随机选着。

		select {
		case v:= <-ch1:
			...
		case v:= <-ch2:
			...
		default:
			...
		}
*/
func main() {

	test_chan()
	test_cache_chan()

}

func test_cache_chan() {

	ch := make(chan int, 10)

	go func() {
		fmt.Println("aa")
		for i := 0; i < cap(ch); i++ {
			rand.Seed(time.Now().UnixNano())
			n := rand.Intn(10)
			fmt.Println("putting : ", n)
			ch <- n
		}

		defer close(ch)
	}()

	fmt.Println("hello from main...")

	for ret := range ch {
		fmt.Println(ret)
	}
}

func test_chan() {
	ch := make(chan int)

	go func() {
		fmt.Println("hello from goroutine ...")
		ch <- 1
	}()

	var ret = <-ch

	fmt.Println("ret: ", ret)
}
