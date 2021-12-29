package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.NewTimer(1 * time.Second)

	defer t.Stop()

	go func() {
		for {
			<-t.C
			fmt.Println("default ...")
			t.Reset(time.Second * 1)
		}
	}()

	time.Sleep(time.Second * 6)
}
