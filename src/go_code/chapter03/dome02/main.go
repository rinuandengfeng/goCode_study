package main

import "fmt"

func main(){
	//golang的变量使用方式1
	//第一种：指定变量的类型，声明后不赋值，使用默认值
	//int的默认值是i =  0
	var i int
	fmt.Println("i = ",i)
	// var f float
	// fmt.Println("f = ",f)
	// string的默认值是空 s =  
	var s string
	fmt.Println("s = ",s)


	//第二种：根据值自行判断变量的类型（类型推导）tow =  10.11
	var tow = 10.11
	fmt.Println("tow = ",tow)

	//第三种：省略var，不能声明已经声明过的变量
	//下面的方式等价与 var name string  name = "tom"
	// :=的 : 不能省略
	name := "tom"
	fmt.Println("name = ",name)
}