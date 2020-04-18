package main
import (
	"fmt"
	"time"
)
type user struct{
	id int
	name string
	city string
	email string
	address string
	creationtime time.Time
	passworld string
}
func main(){
	// 初始化结构体成员
	info := user{
		12,
		"mengc",
		"wuhan","273381276@qq.com",
		"wuhan",
		time.Now(),
		"1076236bb..",
	}
	//初始化部分成员,未初始化的成员默认值为0/时间对象默认值为nil
	/*admin := user{
		id:999,
		name:"admin",
	}
	fmt.Println("struct:",info)
	fmt.Println("struct:",admin)
	//指针变量结构体
	myinfo := &user{
		13,
		"mengc",
		"wuhan","273381276@qq.com",
		"wuhan",
		time.Now(),
		"1076236bb..",
	}
	fmt.Println(myinfo)
	//使用 . 操作符访问结构体下的成员
	fmt.Println(myinfo.id)
	*/
	//newTest(info)
	newTest(&info)
	fmt.Println(info)
}
func newTest(data *user){
	data.id = 666
	fmt.Println(data)
}
