package main

import (
	"fmt"
	"runtime"
)
/*
	并发和并行:
		并行的概念:
			多个任务在同一时间线上同时进行
		并发的概念:
			同一个线程执行多个任务,使用时间片轮转执行多任务,同一时间线上每次只有一个任务执行
	runtime 的函数方法:
		runtime.Gosched() 让出时间片,给子协程执行任务,子协程调用完成时在调用自身
		runtime.Goexit()  退出当前协程
		runtime.GOMAXPROCS() 设置可以并行计算的 CPU 资源中的核心数的最大值,并返回之前的值
	多任务编程:
		资源竞争问题:
			1)
			2)
*/
func main(){
	go func(){
		for i:=0;i<20;i++ {
			fmt.Println("func.go goroutine",i)
		}
	}()
	for i:=0;i<10;i++ {
		runtime.Gosched() //让出时间片,给子协程执行任务,子协程调用完成时在调用自身
		fmt.Println("main.go goroutine",i)
	}
}