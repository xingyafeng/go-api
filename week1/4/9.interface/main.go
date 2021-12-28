package main

import "fmt"

/*

	接口的定义 (多态)

	接口定义了一系列方法的组合
	type IF interface {
		Method(para_list) return_type
	}

	* 适用场景：kubernetes中有大量的接口抽象和多种实现
	* Struct 无需显示申明实现接口interface,只需要实现方法
	* Struct 除实现接口外，还可以实现有额外的方法
	* 一个类型可实现多个结果（Go语言的多重继承）
	* Go语言中的接口不接受属性的定义
	* 接口可以嵌套接口

	注意：
	1、Interface 是可能为nil,所以针对interface的使用一定要预先判空，否则会引起程序crash(nil panic)
	2、Struct 初始化意味着空间的分配，对struct的引用不会出现空指针
*/

type IF interface {
	getName() string
}

type Human struct {
	frist_name, lastname string
}

type Plane struct {
	factory, model string
}

type Car struct {
	factory, model string
}

func (h *Human) getName() string {
	return h.frist_name + "-" + h.lastname
}

func (c *Car) getName() string {
	return c.factory + "-" + c.model
}

func (p *Plane) getName() string {
	return p.factory + "-" + p.model
}

func main() {

	interfaces := []IF{}

	h := new(Human)
	c := new(Car)

	h.frist_name = "first"
	h.lastname = "last"

	interfaces = append(interfaces, h)
	c.factory = "sz"
	c.model = "x5"

	interfaces = append(interfaces, c)

	for _, f := range interfaces {
		fmt.Println(f.getName())
	}

	p := new(Plane)
	p.factory = "sz"
	p.model = "L8"

	fmt.Println("name =", p.getName())
}
