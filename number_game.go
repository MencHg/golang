package main

import (
	"fmt"
	"math/rand"
	"time"
)
func createNumber() (count int) {
	rand.Seed(time.Now().UnixNano())
	for{
		count = rand.Intn(1000000)
		if count > 100000{
			break
		}
	}
	//fmt.Println("count",count)
	return
}
func main(){
	number := createNumber()
	fmt.Println(number)
}
