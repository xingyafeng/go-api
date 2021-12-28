package main

import (
	"encoding/json"
)

/*
	json <-> struct

	Unmarshal string -> struct
	Marshal   struct -> string

*/

type Stu struct {
	Name  string `json:"name"`
	Age   int
	HIgh  bool
	sex   string
	Class *Class `json:"class"`
}
type Class struct {
	Name  string
	Grade int
}

// func unmarshal2struct(humanStr string) Human {
// 	h := Human{}

// 	err := json.Unmarshal([]byte(humanStr), h)
// 	if err != nil {
// 		println(err)
// 	}

// 	return h
// }

func marsha2JsonString(s Stu) string {

	updateBytes, err := json.Marshal(s)
	if err != nil {
		println(err)
	}

	return string(updateBytes)
}

/*
	注意事项：
	1、只要是可导出成员（变量首字母大写），都可以转成json。因成员变量sex是不可导出的，故无法转成json。
	2、如果变量打上了json标签，如Name旁边的 `json:"name"` ，那么转化成的json key就用该标签“name”，否则取变量名作为key，如“Age”，“HIgh”。
		bool类型也是可以直接转换为json的value值。Channel， complex 以及函数不能被编码json字符串。当然，循环的数据结构也不行，它会导致marshal陷入死循环。
	3、指针变量，编码时自动转换为它所指向的值，如cla变量。
	4、（当然，不传指针，Stu struct的成员Class如果换成Class struct类型，效果也是一模一样的。只不过指针更快，且能节省内存空间。）
	5、最后，强调一句：json编码成字符串后就是纯粹的字符串了。

	https://blog.csdn.net/zxy_666/article/details/80173288

	json 编解码
*/

func main() {

	//实例化一个数据结构，用于生成json字符串
	stu := Stu{
		Name: "张三",
		Age:  18,
		HIgh: true,
		sex:  "男",
	}

	//指针变量
	cla := new(Class)
	cla.Name = "1班"
	cla.Grade = 3
	stu.Class = cla
	println(marsha2JsonString(stu))
}
