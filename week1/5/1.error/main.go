package main

import (
	"errors"
	"fmt"
)

/*
	error 构建 需要扩展，如：需要打印具体的函数，哪一行出现的错误。

	构建方法：
	1、fmt.Errorf(string)
	2、errors.New(string)

	注意事项：
	1、当我们在 Golang 中使用 errors.New("Aaa.") 形式返回 error 信息时，文字内容不应该以大写字母开头或者标点符号结尾。
*/

type StatusError struct {
	ErrStatus string
}

func (e *StatusError) Error() string {

	return e.ErrStatus
}

func main() {

	// 1.
	err := fmt.Errorf("it is err")
	fmt.Println("err:", err)

	// 2.
	var err1 error = errors.New("NotFound")
	fmt.Println("err1:", err1)

	var _error = &StatusError{"error"}
	fmt.Println("errx: ", _error)
}
