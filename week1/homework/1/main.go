package main

import "fmt"

func main() {
	// 数组的定义
	arr := []string{"I", "am", "stupid", "and", "weak"}

	fmt.Println("before arr is :", arr)

	for i, v := range arr {
		switch v {
		case "stupid":
			arr[i] = "smart"
		case "weak":
			arr[i] = "strong"
		}
	}

	fmt.Println("after  arr is :", arr)
}
