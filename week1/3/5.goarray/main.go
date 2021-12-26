package main

import "fmt"

func main() {

	array := [3]int{1, 2, 3}

	fmt.Println("array = ", array)

	for _, arr := range array {
		println(arr)
	}
}
