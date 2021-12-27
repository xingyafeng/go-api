package main

// 类型的别名定义
type ServiceType string

const (
	ServiceType1 ServiceType = "one"
	ServiceType2 ServiceType = "two"
	ServiceType3 ServiceType = "three"
)

func main() {

	println(ServiceType1)
	println(ServiceType2)
	println(ServiceType3)
}
