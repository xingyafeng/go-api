package main

import (
	"fmt"

	_ "github.com/xingyafeng/goapi/week1/4/2.init/a"
	_ "github.com/xingyafeng/goapi/week1/4/2.init/b"
)

func init() {

	fmt.Println("#init : main")
}

// 导入本地包名的时候，需要注意：
// 1. import name 对应包中
// 2. 导包的顺序与main中一致，当存在多次导包，相同的只导一次
// 3. init 函数执行顺序是压栈的方式，先进后出原则
func main() {

	print("test start ...")

}
