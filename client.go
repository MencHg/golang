package main

import (
	"fmt"
	"net"
)

/*

*/
func main(){
	clientCtx,err := net.Dial("tcp","127.0.0.1:8080")
	if err!=nil{
		fmt.Println("Error:",err)
		return
	}
	defer clientCtx.Close()
	clientCtx.Write([]byte("message"))

}