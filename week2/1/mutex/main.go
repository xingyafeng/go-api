package main

import (
	"fmt"
	"sync"
	"time"
)

func rLock() {

	lock := sync.RWMutex{}

	go func() {
		for i := 0; i < 3; i++ {
			lock.RLock()

			defer lock.RUnlock()
			fmt.Println("rLock:", i)
		}
	}()
}

// 问题: GO loop variable captured by func literal
// 解决方法：
// 1. 闭包，且传入参数
// 2. 重新声明
// 3. 多写一层调用
func wLock() {
	lock := sync.RWMutex{}

	for i := 0; i < 3; i++ {
		// go func(i int) {
		fmt.Println(i)
		lock.Lock()
		defer lock.Unlock() // 注意：这个解锁过程是在当前代码块退出就执行的
		fmt.Println("wlock:", i)
		// }(i)
	}
}

func lock() {
	lock := sync.Mutex{}
	for i := 0; i < 3; i++ {
		// go func(i int) {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("lock:", i)
		// }(i)
	}
}

// 加锁
// 1. sync.Mutex
// 2. sync.RWMutex
func main() {

	go rLock()
	go wLock()
	go lock()

	time.Sleep(time.Second * 5)
}
