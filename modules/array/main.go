package main

import "fmt"

func main() {

	// 数组的申明
	var arr [10]int
	arrx := [10]int{1, 2, 3, 4}

	for i := 0; i < len(arr); i++ {
		arr[i] = i + 3 //赋值
	}

	for i, v := range arr {
		fmt.Println("i = ", i, "; v = ", v)
	}

	for _, v := range arrx {

		fmt.Println("v = ", v)
	}

	// fmt.Println("main")
}
