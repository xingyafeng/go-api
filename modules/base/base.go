package main

import (
	"fmt"
)

func test_if(x int) int {
	var v int

	if v := x - 100; v < 0 {
		return v
	}

	return v
}

func test_switch() {

	var var1 string = "tct"

	// fmt.Println("test switch ...")
	switch var1 {
	case "tct":
	case var1:
		fallthrough //向下穿透
	case "test":
		fmt.Println("test ...")
	default:
		fmt.Println("default ...")
	}
}

func test_for() {

	var count int = 100
	var sum int
	//1
	for i := 0; i < count; i++ {
		sum += i
	}
	fmt.Println("sum: ", sum)

	//2
	for sum < count {
		sum = +sum
	}

	fmt.Println("sum 2: ", sum)
	//3
	for {
		if sum > 100 {
			fmt.Println("sunm 3: ", sum)
			break

		} else {
			fmt.Println("else ...")
		}

	}

	// one string
	var str string = "tct"
	for index, value := range str {
		fmt.Println("# ", index, " ->", string(value))
	}
	fmt.Println("str = ", str)

	// two array
	var arr [5]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}

	for _, value := range arr {
		fmt.Println("value = ", value)
	}

	// three slice
	var ss []int
	fmt.Println("ss = ", ss, len(ss), cap(ss))

	for i := 0; i < 21; i++ {
		ss = append(ss, i+1)
	}
	fmt.Println("ss = ", ss, len(ss), cap(ss))

	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}

}

func main() {

	// 1、if
	fmt.Println(test_if(100))

	// 2、switch
	test_switch()

	// 3、for
	test_for()
	ret, error := fmt.Println("a")
	fmt.Println(ret, error)
}
