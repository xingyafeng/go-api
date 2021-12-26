package main

import "fmt"

func deleteItem(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
func main() {

	myArray := [5]int{1, 2, 3, 4, 5}
	mySlice := myArray[2:5]
	fullSlice := myArray[:]

	fmt.Println("mySlice: ", mySlice)
	remove3item := deleteItem(fullSlice, 3)
	fmt.Println("del: ", remove3item)
}
