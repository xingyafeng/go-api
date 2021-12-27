package main

import "fmt"

/*
	指针

	程序在内存中存储它的值，每个内存块（或字）有一个地址，通常用十六进制数表示，如：0x6b0820 或 0xf84001d7f0。

	Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。

	var intP *int
	intP = &i1

	intP 存储了 i1 的内存地址；它指向了 i1 的位置，它引用了变量 i1

	注意：
	符号 * 可以放在一个指针前，如 *intP，那么它将得到这个指针指向地址上所存储的值；这被称为反引用（或者内容或者间接引用）操作符；另一种说法是指针转移。
	对于任何一个变量 var， 如下表达式都是正确的：var == *(&var)。

*/

func test_pointer_one() {
	var i = 5
	var name string = "s"

	var intP *int = nil
	var intS *string = nil

	intP = &i
	intS = &name

	fmt.Printf("i= %d, its local in memory :%p\n", i, &i)

	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
	fmt.Printf("The value at memory location %p is %s\n", intS, *intS)
}

// 特别注意:
// 1. 指针复制非常危险
// 2. 不能得到一个文字或常量的地址
func test_pointer_two() {

	var s string = "string"
	var p *string = &s

	fmt.Printf("s = %s\n", s)

	// 修改指针的内容
	*p = "xxx"

	fmt.Printf("p = %s, s = %s\n", *p, s)
}

// 对一个空指针的反向引用是不合法的，并且会使程序崩溃
func test_pointer_three() {

	var s = 1
	var p *int = nil
	// p = &s
	*p = 2
	fmt.Printf("p = %d, %d, %p\n", s, *p, p)

}
func main() {

	test_pointer_one()
	test_pointer_two()
	test_pointer_three()

}
