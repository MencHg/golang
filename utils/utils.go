package utils
import "fmt"
func Utils(count,page int){
	fmt.Println("这是第:",page,"页 的第",count,"条数据")
}
func Kit() (a,b,v int) {
	a,b,v = 1,2,3
	return
}
func MathRand(count,page int) (max,min int){

	if count>page {
		max,min = count,page
	}else {
		max,min = page,count
	}
	return
}