package main

import "fmt"

/*

	context <https://zhuanlan.zhihu.com/p/110085652>

	为什么需要context?
		* 在并发程序中，由于超时、取消操作或者一些异常情况，往往需要进行抢占操作或者中断后续操作。
		* 熟悉channel的朋友应该都见过使用done channel来处理此类问题。
*/
func main() {

	fmt.Println("context")
}
