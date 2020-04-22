package main

import "net"

func main(){
	clientServer()
}
func clientServer(){
/*
	clientServer:
		code...
*/
	//主动链接服务器
	client,err := net.Dial("tcp","127.0.0.1:8080")
	if err != nil{
		return
	}
	defer client.Close()
	client.Write([]byte("9527"))
}