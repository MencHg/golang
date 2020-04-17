package main

import "fmt"

func main(){
	/*
		切片的特点:
			1.和数组的区别,数组的[] 符中的长度规定是一个常量,数组不能改变长度,len和cap永远都是定义时的值
			2.切片[]/[...] 符中可以为空或者为... 操作符,长度可以不固定
	*/
	array_ := [6]int {0,1,2,3,4,5}
	slice_ := array_[0:3:5]
	fmt.Println("main 函数:slice",slice_)
}