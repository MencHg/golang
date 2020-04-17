package main

import (
	"fmt"
)

func swap(count,page int){ //函数接收参数
	count,page = page,count // 值交换
	fmt.Println(count,page) //打印交换后的结果 输出交换后的值 count,page 交换后 20,10
}
func main(){
	count,page := 10,20
	//swap(count,page) // 值传递
	quest(&count,&page) //传递变量的内存地址
	fmt.Println(count,page) //这里的count,page 还是原来的的值10,20
}
func quest(count,page *int){ //这里传递的是变量的内存地址
	*count,*page = *page,*count // 同一变量的内存地址发生改变,内存地址与传递前的变量内存地址一样,所以会使传递前的原始值发生改变
	fmt.Println(*count,*page) //此时打印结果和主函数的打印结果一直,
}