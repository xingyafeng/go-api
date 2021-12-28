package main

import (
	"fmt"
	"reflect"
)

/*

	反射机制
	1、reflect.TypeOf()  返回被检查对象的类型
	2、reflect.VauleOf() 返回被检查对象的值
*/
func test_reflect() {
	myMap := make(map[string]string, 10)

	myMap["a"] = "b"

	t := reflect.TypeOf(myMap)
	fmt.Println("type: ", t)

	v := reflect.ValueOf(myMap)
	fmt.Println("value : ", v)

	fmt.Println("srcstr: ", myMap)

}

// 反射的方法
func main() {

	// 测试基础的类型的反射
	test_reflect()

	myStruct := T{A: "a"}

	v := reflect.ValueOf(myStruct)

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Filed: %d:%v\n", i, v.Field(i))
	}

	for i := 0; i < v.NumMethod(); i++ {
		fmt.Printf("Method: %d:%v\n", i, v.Method(i))
	}

	// 返回内容是数组
	result := v.Method(0).Call(nil)

	fmt.Println("result = ", result[0])
	fmt.Println("result = ", myStruct.Test())
}

type T struct {
	A string
}

// 需要注意receiver是值还是指针？
func (t T) Test() string {
	return t.A + "1"
}
