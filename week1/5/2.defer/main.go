package main

import (
	"fmt"
	"sync"
	"time"
)

func loopFunc() {

	lock := sync.Mutex{}

	for i := 0; i < 3; i++ {
		func() { // 闭包解决加锁的问题。 defer在函数退出之前退出
			lock.Lock()
			defer lock.Unlock()

			fmt.Println("loopFunc:", i)
		}()
	}
}

/*
	defer 滞后 推后执行, 在程序退出之前做的一些事情。

	原理：defer 按照压栈方式 ，打印是pop stack

	用途：
		1、在函数退出之前执行某个语句或者函数
		2、场景的defer使用场景：记得关闭你打开的资源
			* defer file.CLose()
			* defer mu.Unlock()
			* defer println("ddddd")

	panic：可在系统出现不可恢复错误时主动调用panic ,panic会使当前线程直接crash
	defer: 保证执行并把控制权交还给接受到panic的函数调用者
	recover: 函数从panic或错误场景中恢复
*/
func main() {

	defer fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")

	loopFunc()

	defer func() {
		fmt.Println("defer func is called")

		if err := recover(); err != nil {
			fmt.Println("err:", err)
		}
	}()

	time.Sleep(time.Second)

	panic("a panic is triggered ...")
}
