package main

import (
	"fmt"
	"time"
)

var ret int

func main() {
	ret = 100
	fmt.Println("<00> ret:", ret)
	go test_thread1()
	go test_thread2()
	fmt.Println("<01> ret:", ret)

	for {
		// time.Sleep(time.Second * 1)
		println("ret :", ret)
	}
}

func test_thread1() {

	for {

		ret = 101

		fmt.Println("<1> ret:", ret)
		time.Sleep(time.Second)
	}
}

func test_thread2() {

	for {

		ret = 102

		fmt.Println("<2> ret: ", ret)
		time.Sleep(time.Second * 5)
	}
}
