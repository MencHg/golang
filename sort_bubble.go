package main

import (
	"fmt"
	"math/rand"
	"time"
)
func main(){
	slice_ := []int{}
	slice_ = SlicePush()
	fmt.Println("slice_nosort:",slice_)
	BubbleSort(slice_)
	fmt.Println("slice_ensort:",slice_)
}
func BubbleSort(value []int){  //冒泡排序函数
	valueLength := len(value)
	for i:=0;i<valueLength-1;i++{
		for j:=0;j<valueLength-1-i;j++{
			if value[j]>value[j+1] {
				value[j],value[j+1] = value[j+1],value[j]
			}
		}
	}
}
func SlicePush()(value []int){
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<20;i++{
		value = append(value,rand.Intn(100))
	}
	//fmt.Println(value)//打印执行结果
	return //返回添加数据后的切片
}