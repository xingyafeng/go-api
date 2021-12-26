package main

import (
	"fmt"
	"time"
)

// 纤尘 需要抓住时间片才能运行否则无法执行的？
func main() {

	ret := 1

	println('a')

	go func(ret int) {
		fmt.Println(ret)
	}(ret)

	go fmt.Println("2")

	time.Sleep(time.Second)
}
