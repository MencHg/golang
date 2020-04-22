package main

import "fmt"

func main(){
	heroList := []string{"宋江","林冲","鲁智深","武松"}
	for index,ele := range heroList{
		fmt.Println("这是第",index,"个名为",ele,"的英雄\n")
	}
}