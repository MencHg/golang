package main

import "fmt"

func main(){
	arr := [10]int {0,1,2,3,4,5,6,7,8,9}
	test(arr) //数组作为参数传递给函数调用时,是将原数组copy 一份给调用方
	test2(&arr) //使用数组指针作为参数传递给函数时,是将原数组的内存地址给调用函数
	fmt.Println("main 函数",arr)
}
func test(arr [10]int){
	arr[1] = 10 //这里改变传递过来的数组,不会影响到原数组
	fmt.Println("test函数",arr)
}
func test2(arr_ *[10]int){
	(*arr_)[0] = 99 //这里改变传递过来的数组,会同时改变调用前的原数组
	fmt.Println("test2函数",*arr_)
}