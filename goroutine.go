package main

import (
	"fmt"
	"time"
)
/*
	goroutine 特性:
		1.创建一个新的协程 使用 go 关键字调用函数
		2.主携程先调用完成跳出时,子协程也将终止执行
*/
func main (){
	go newTask() //代码示例
	for  i:=0;i<15;i++{
		fmt.Println("No.",i,"this is newTask goroutine")
		time.Sleep(time.Second)
		if i==15{
			break //主协程调用完成跳出时子协程也将终止
		}
	}
}
func newTask(){
	for  i:=0;i<100;i++{
		fmt.Println("No.",i,"this is main goroutine")
		time.Sleep(time.Second)
	}
}