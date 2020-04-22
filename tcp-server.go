package main

import (
	"fmt"
	"net"
)
/*
	tcp-server
*/
func main(){
	tcpServer() //调用 tcpServer 服务函数
}
func tcpServer(){
	fmt.Println("TpcServer starting to 8080 port ....")
	httpListen,err := net.Listen("tcp","127.0.0.1:8080")
	if err!=nil{
		fmt.Println("Error:HttpListen",err)
		return
	}
	defer httpListen.Close()
	// 阻塞等待用户链接
	httpCtx,err_ := httpListen.Accept()
	if err_ != nil{
		fmt.Println("Error: Accept",err_)
		return
	}
	buffer_ := make([]byte ,1024)
	data,error_ := httpCtx.Read(buffer_)
	if error_ != nil{
		fmt.Println("Error:Buffer",error_)
		return
	}
	fmt.Println("data:",string(data))
	defer httpCtx.Close()
}