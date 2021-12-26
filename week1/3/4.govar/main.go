package main

func test_var() {

	// 显示转换
	var a int = 100
	var f float32 = (float32(100))
	var u uint = uint(f)

	// 推导
	var i int
	j := i

	println(a, f, u, i, j)
}

func main() {
	// 1. const 常量 不可修改
	const a = "hello"

	// 2. 变量
	var ret int

	var i, j, k = 0, 1, 2
	c, python, java := 1, 's', false

	//1
	print(a)
	//2
	print(ret)

	print(i, j, k)
	print(c, python, java)

}
