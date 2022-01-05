package main

import (
	"fmt"
	"net/http"
)

func myResponse(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<html><center> <font size=\"40\">Hello！ I am go service started by Alexander!</font></center></html>"))

	fmt.Println("w.Header()", w.Header())

	fmt.Println("r.Header", r.Header)
	fmt.Println("A client has just visited!")
}

/*
编写一个 HTTP 服务器，大家视个人不同情况决定完成到哪个环节，但尽量把 1 都做完：

    1; 接收客户端 request，并将 request 中带的 header 写入 response header
    2; 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
    3; Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
    4; 当访问 localhost/healthz 时，应返回 200
*/
func main() {

	http.HandleFunc("/", myResponse)
	http.ListenAndServe("127.0.0.1:8888", nil)
}
