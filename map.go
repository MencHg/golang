package main
import "fmt"
func main(){
	//定义map
	map_ := map[string]string{
		"React":"v_1.17.1",
		"Nginx":"v_1.0.01",
		"Vuejs":"v_2.6.1",
		"MongoDB":"v_4.1.0",
	}
	maple_ := map[int]string{
		0:"nginx",
		1:"react",
		2:"vuejs",
		3:"mongodb",
	}
	fmt.Println(map_,"\n",maple_)
	//通过访问键名key,修改对应键值value
	map_["Nginx"] = "v_2.16.1"
	maple_[3] = "mongoose"
	//如果对应的键值不在访问的map中则会自动扩容,创建对应的key:value
	maple_[4] = "js"
	map_["xiaoming"] = "v_2.23"
	fmt.Println(map_,"\n",maple_)
	//map的遍历
	for key,value := range maple_{
		fmt.Println(key,value)
	}
	mapIskey(maple_,2)
	fmt.Println(maple_)
}
func mapIskey(data map[int]string,index int){
	value,iskey := data[index]
	//map的删除操作 delete(目标map对象,需要删除操作的key)
	delete(data,0)
	//判断目标中是否存在该key
	if iskey{
		fmt.Println("data:",value)
	}else {
		fmt.Println("访问的目标中不存在该key")
	}
}