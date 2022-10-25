package main

import "fmt"

//定义全局变量
var n1 = 100
var n2 = 200
var name = "jack"
//一次性声明多个全局变量
var (
	n3 = 300
	n4 = 900
	name1 = "mary"
)


func main(){
	//输出全局变量 n1 =  100 n2 =  200 name =  jack
	fmt.Println("n1 = ",n1, "n2 = ",n2, "name = ",name)

	//数出一次性定义的多个全局变量  n1 =  100 n2 =  200 name =  jack
	fmt.Println("n3 = ",n3, "n4 = ",n4, "name1 = ",name1)
}