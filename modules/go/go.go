package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)

	ticker := time.NewTicker(time.Second)
	go func() {
		for i := 0; i < 3; i++ {
			<-ticker.C
			ch <- i
		}
		defer close(ch)
		defer ticker.Stop()
	}()

	for ret := range ch {
		fmt.Println(ret)
	}
	test_interface()
}

func test_interface() {

	type test struct{}

	v := test{}

	printv(v)
}

func printv(v interface{}) {

	println(v)
}
