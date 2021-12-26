package main

import (
	"fmt"
)

func main() {

	// var map[string]int

	myMap := make(map[string]int, 10)

	myMap["a"] = 10

	myFuncMap := map[string]func() int{
		"funcA": func() int {
			return 100
		},
	}

	//1
	fmt.Println("myMap: ", myMap)

	//2
	fmt.Println(myFuncMap)

	f := myFuncMap["funcA"]
	fmt.Println("f:", f())

	// 3 取出map中的值
	vaule, exist := myMap["a"]
	if exist {
		println(vaule)
	}

	// 遍历其中的值
	for key, value := range myMap {
		println(key, value)
	}
}
