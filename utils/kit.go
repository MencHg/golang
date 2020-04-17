package utils

import "fmt"

func Sdk(){
	store := 0
	func(){
		store++
		fmt.Println("这是sdk函数:",store)
	}()
}