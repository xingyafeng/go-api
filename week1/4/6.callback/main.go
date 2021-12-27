package main

import "fmt"

func DoOperation(y int, f func(int, int)) {
	f(y, 1)
}

func inc(a, b int) {
	fmt.Printf("@ a+b=%d\n", a+b)
}

func dec(a, b int) {
	fmt.Printf("@ a-b=%d\n", a-b)
}

// callback 把函数名作为参数
func main() {
	DoOperation(1, inc)
	DoOperation(1, dec)
}
