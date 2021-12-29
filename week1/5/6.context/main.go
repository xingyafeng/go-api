package main

import (
	"context"
	"fmt"
	"time"
)

/*

	context <https://zhuanlan.zhihu.com/p/110085652>

	为什么需要context?
		* 在并发程序中，由于超时、取消操作或者一些异常情况，往往需要进行抢占操作或者中断后续操作。
		* 熟悉channel的朋友应该都见过使用done channel来处理此类问题。

	用法：
		1、context.Background
			Background通常被用于主函数、初始化以及测试中，作为一个顶层的context,一遍创建context的时候都是Background
		2、context.TODO
			TODO是在你不确定使用那个context的时候才会使用
		3、context.WithDeadline
			超时时间
		4、context.WithValue
			想context 添加键值对
		5、context.WithCancel
			创建一个可取消的context

	学习：
		1、创建context方法 context.Background
		2、context 在线程中的重要地位
		3、如何使用for range ticker.C {} 每秒执行一次

*/
func main() {

	baseCtx := context.Background()
	ctx := context.WithValue(baseCtx, "key", "value")

	go func(c context.Context) {
		if v := c.Value("key"); v != nil {
			fmt.Println(c.Value("key"))
		}
	}(ctx)

	time.Sleep(time.Second)

	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	go func(ctx context.Context) {
		ticker := time.NewTicker(time.Second)

		for range ticker.C { // 每秒执行一次for循环，相当于延时1s
			select {
			case <-ctx.Done():
				fmt.Println("child process interrupt ...")
				return
			default:
				fmt.Println("enter child process default...")
			}
		}
	}(timeoutCtx)

	select {
	case <-timeoutCtx.Done():
		time.Sleep(time.Second)
		fmt.Println("main process exit ...")
	default:
		fmt.Println("enter main default ...")
	}

	// 保证栗子可以正常工作 延时3s 不然 子线程不会超时退出，进入死循环。
	time.Sleep(time.Second * 3)
}
