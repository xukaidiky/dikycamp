package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

/**
	main入口
**/
func main() {
	//1、注册handle处理函数healthz，声明用户访问http服务路径为"/"
	http.HandleFunc("/", healthz)
	//2、ListenAndService绑定，给healthz绑定8080端口，地址ip为0.0.0.0
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

/**
	handle处理函数healthz的定义
**/
func healthz(w http.ResponseWriter, r *http.Request) {
	//work1、接收客户端request，并将request中带的header写入response header
	//判断客户端request是否有长度，不能为空
	if len(r.Header) > 0 {
		//通过for..range获取request的Header信息
		for k, v := range r.Header {
			//每次循环将request的Header信息写入响应response的Header中
			log.Println("%s=%s", k, v[0])
			w.Header().Set(k, v[0])
		}
	}

	//work2、读取当前系统的环境变量中的VERSION配置，并写入response header
	//设置一个VERSION的环境变量为go.1.20.5
	os.Setenv("VERSION", "go version 1.20.5")
	//获取刚设置的VERSION的环境变量
	versionName := os.Getenv("VERSION")
	//读取VERSION
	log.Println("VERSION env is", versionName)
	w.Write([]byte(versionName + "\n"))

	//work3、Server端记录访问日志包括客户端IP，HTTP返回码，输出到server端的标准输出
	//获取request请求客户端的IP
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Fatal("err:", err)
	}
	//对客户端IP信息判空并进行标准输出
	if net.ParseIP(ip) != nil {
		fmt.Printf("Client ip addr is: %s \n", ip)
		log.Println("Client ip addr is: ", ip)
	}
	//获取HTTP返回码
	statusCode := http.StatusOK
	//对HTTP返回码进行标准输出
	fmt.Printf("HTTP Status Code is: %d \n", statusCode)
	log.Println("HTTP Status Code is: ", statusCode)

	//work4、当访问localhost/healthz时，应返回200
	w.WriteHeader(statusCode)
	w.Write([]byte("HTTP Server is Suceess!"))
}
