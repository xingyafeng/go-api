package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")

	var scores []int

	fmt.Println(len(scores), cap(scores))

	s := scores[0:9]
	s[0] = 1
	// ss := append(scores, 11)

	fmt.Println("s: ", s)

	// asdfasdf

}
