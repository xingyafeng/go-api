package main

import (
	"fmt"
)

/*
	闭包 closure 匿名函数

		* 不能独立存在
		* 可以赋值给其他变量 x:= func(){}
		* 可以直接调用
			func(x,y int) {println(x+y)(1,2)}
		* 可作为函数的返回值
			func add()(func(b int) int)

	使用场景
		defer
*/
func main() {

	func() {
		fmt.Println("hello closure")
	}()

	defer func() {
		if r := recover(); r != nil {
			println("recovered in FunX ...")
		}
	}()
}
