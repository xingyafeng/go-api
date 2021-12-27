package main

import "fmt"

/*
	可变参数

	Go语言中的可变长参数允许调用方传递任意多个相同类型的参数

*/

func delete(slice []int, index int) []int {

	return append(slice[:index], slice[index+1:]...)
}

func main() {

	arr := [5]int{1, 2, 3, 4, 5}
	ss := arr[:]

	fmt.Println(ss)
	ss = delete(ss, 2)
	fmt.Println("delete: ", ss)
	fmt.Println(ss)
	fmt.Println("new array: ", append(ss, 100))
}
