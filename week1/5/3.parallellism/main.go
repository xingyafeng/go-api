package main

import (
	"fmt"
	"time"
)

/*

	并发与并行
	并发：多个线程在一个cpu上执行 （两个或多个事件在同一时间间隔发生）
	并行：计算机有多个cpu，不同线程在不同的cpu上执行。（两个或多个事件在同一时刻方式）

	协程

		* 进程： （内核态） 切换都是一个个的task
			* 分配系统资源（CPU时间，内存等）基本单位
			* 有独立的内存空间，切换开销大 （多大呢？与线程和协程对比数据是什么样呢？）

		* 线程: 进程的一个执行流，是CPU调度并能执行的独立运行的基本单位  （内核态）
			* 同一个进程中的多线程共享内存空间，线程切换代价小 （线程的切换不切地址空间）
			* 多线程通信方便
			* 从内核层面来看线程其实是一种特殊的进程，他跟父进程共享了打开的文件和文件系统信息，共享了地址空间和信号处理函数

		* 协程 （用户态） 抓住时间片不放出去。
			* Go语言中轻量级线程的实现
			* Golang 在runtime 、系统调用等多方面对goroutine调度进行了封装和处理，当遇到时间长执行或者进行系统调用时，
			  会主动把当前的goroutine的CPU(p)转让出去，让其他的goroutine 能被调度并执行，也就是golang从语言层面支持了协程。


	CSP模型 （Communicating Sequential Process）

		* csp
			描述两个独立的并非实体通过共享的通信channel进行通信的并非模型。

		* Go协程gorountine
			* 是一种轻量级的线程，它不是操作系统的线程，而是将一个操作系统的线程分段使用，通过调度器实现协作式调度。
			* 是一中绿色线程，微线程，他与Coroutine 协程也有区别，能够在发现阻塞的后启动新的微线程。

		* 通道channel
			* 类似Unix 的pipe ,用于协程的之间的通信与同步。 协程之间虽然解耦，但他们和Channel有着耦合。

	线程与协程的差异对比

		* 每个goroutine 默认占用内存远比java、C++的线程少 （少多少呢？）
			* goroutine : 2KB
			* 线程： 2M 8M

		* 线程goroutine 切换开销方面，goroutine 远比线程小

			* 线程： 涉及模式切换（从用户态陷入内核态）、16个寄存器、PC、SP.. 等寄存器的刷新
			* goroutine : 只有三个寄存器的修改，- PC/SP/DX

		* GOMAXPROCS
			* 控制并行的线程数量
*/

func main() {

	go fmt.Println("a")
	go fmt.Println("b")
	go fmt.Println("c")

	time.Sleep(time.Second)
}
