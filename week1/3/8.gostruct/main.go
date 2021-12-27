package main

import (
	"fmt"
	"reflect"
)

type person struct {
	name string
	id   int
}

type T struct {
	Name string `json:"tct"`
}

/*
	结构体的内存布局

	Go 语言中，结构体和它所包含的数据在内存中是以连续块的形式存在的，即使结构体中嵌套有其他的结构体，这在性能上带来了很大的优势。
不像 Java 中的引用类型，一个对象和它里面包含的对象可能会在不同的内存空间中，这点和 Go 语言中的指针很像。

	递归结构体

	结构体类型可以通过引用自身来定义。这在定义链表或二叉树的元素（通常叫节点）时特别有用，此时节点包含指向临近节点的链接（地址）。
	如下所示，链表中的 su，树中的 ri 和 le 分别是指向别的节点的指针。

	结构体转换

*/

func test_base() {
	// t是指向一个结构体类型变量的指针
	var p person

	var t *person

	// p是结构体类型变量 方法一
	p = person{
		name: "tct",
		id:   2,
	}

	// 无需定义 使用 ':=' 方法二
	p1 := person{
		name: "tct",
		id:   2,
	}

	// 初始化方法一
	t = new(person)
	t.name = "tct"
	t.id = 1

	// 初始化方法二
	t1 := &person{}
	t1.name = "tct"
	t1.id = 1

	// 常用的初始化方法
	var ps person
	ps.name = "tct"
	ps.id = 100

	fmt.Println(ps)

	fmt.Println(p)
	fmt.Println(p1)
	fmt.Println(t)
	fmt.Printf("p %v\n", p)
	fmt.Printf("t %v\n", t)
	fmt.Printf("t1 %v\n", t1)
}

func test_tag() {

	var t T
	t.Name = "tct"

	// 反射拿到类型和值
	tag := reflect.TypeOf(t).Field(0).Tag.Get("json")

	print(tag)
}

func main() {
	test_base()
	test_tag()
}
