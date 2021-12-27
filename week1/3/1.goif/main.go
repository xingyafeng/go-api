package main

// if statement; condition {
// }

func test_if(x int) int {
	var v int

	if v := x - 100; v < 10 {
		return v
	}

	return v
}

func main() {

	println(test_if(1))
}
