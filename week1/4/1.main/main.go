package main

import (
	"flag"
	"fmt"
	"os"
)

// main 函数的入参方法：
// 1. os.Args
// 2. flag.String 支持不同类型的参数
//    flag.Parse() 使用解析则使用命令行的参数，否则使用默认
func main() {

	fmt.Println("args: ", os.Args[1:])

	name := flag.String("name", "tct", "this is name key")
	flag.Parse()

	println(*name)
}
