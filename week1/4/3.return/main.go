package main

import (
	"errors"
	"fmt"
)

func test_return(x, y int) (int, error) {

	if v := x; v < 0 && y < 0 {
		return 0, errors.New("math. is error")
	}

	return x + y, nil
}

func pass_value(x, y int) (a, b int) {
	a = x
	b = y

	return a, b
}

/*
好处：
	1. 多返回值 ---- 任意数量
	2. 命名返回值
		* Go的返回值可命名，他们可视作定义在函数顶部的变量。
		* 返回值的名称具有一定的意义，它可作为文档使用
		* 没有参数的return语句返回已命名的返回值。就是直接返回
	3. 可或略其中某个返回值，是用占位符 ’_‘

**/
func main() {

	result, err := test_return(1, 2)

	if err != nil {
		fmt.Println("result: ", result, err)
	}

	fmt.Println("@@ result: ", result)

	a, _ := pass_value(100, 2)
	fmt.Println("a = ", a)
}
