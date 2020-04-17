package main
import (
	"fmt"
	"./utils"
)
func main(){
	str := "123"
	utils.Utils(20,1)
	key,value,count := utils.Kit()
	k,v := utils.MathRand(10,8)
	fmt.Println(str,":",key,value,count)
	fmt.Println(k,v)
	utils.Sdk()
}