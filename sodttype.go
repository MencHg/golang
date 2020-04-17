package main

import "fmt"

func main(){
	a := 25
	fmt.Printf("a = %d\n",a) //输出变量内存
	fmt.Printf("&a = %p\n",&a) // 输出变量地址
	var p *int
	p = &a // 这里改变的p的值,也改变a的值 因为内存地址一样
	fmt.Printf("p = %p\n",p)
	*p = 666	//*p 操作的不是p的内存,是p所指向的内存地址(也就是a)
	fmt.Printf("*p = %p\n",*p,a) //输出*p 和 a
}