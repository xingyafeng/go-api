package main

import "fmt"

func test_switch() {

	var var1 string = "tct"

	fmt.Println("test switch ...")
	switch var1 {
	case "tct":
		fmt.Println("tct ...")
		fallthrough //向下穿透
	case "test":
		fmt.Println("test ...")
	case var1:
	default:
		fmt.Println("default ...")
	}
}

func main() {

	test_switch()
}
