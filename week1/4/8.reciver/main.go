package main

import "fmt"

/*
	方法
	作用在接受者上的函数

	使用场景
		很多场景下，函数需要的上下文可以在receiver属性中，通过定义receiver的方法，
		该方法可以直接访问receiver属性，减少参数传递的需求

	传值还是传指针
		Go语言只有一种规则：传值
		函数内修改参数的值不会影响函数外原始变量的值
		可以传指针参数将变量地址传递给调用函数，Go语言会
		复制该指针的地址作为函数内的地址，但是指向统一地址。
*/

type Person struct {
	name string
	id   int
}

func (p *Person) getPreson() Person {
	return *p
}

func main() {

	var p Person
	p.name = "tct"
	p.id = 1

	fmt.Printf("name: %s\n ", p.getPreson().name)
	fmt.Printf("id: %d\n", p.getPreson().id)
}
