package main

import (
	"fmt"
	"time"
)

type User interface { //定义接口
	createuser()
	uptdate()
}
type Person struct {
	Id           int
	Name         string
	Like         []string
	address      string
	Birthday     string
	Creationtime time.Time
}
func (data *Person) createuser(){
	fmt.Println(data.Name,data.address)
}
func (data *Person) uptdate(){
	fmt.Println(data.Name,data.address)
}
func TestFunction(args interface{}){
	fmt.Println(args)
}
func main(){
	fmt.Println("Person",)
}