package main

import "fmt"

func main(){

	//演示golang如何一次性定义多个变量
	// 输出结果n1 =  0 n2 =  0 n3 =  0
	var n1, n2, n3 int
	fmt.Println("n1 = ",n1, "n2 = ",n2, "n3 = ",n3)

	//一次性声明多个变量的方式2
	//输出结果为 n4 =  100 name =  tom n5 =  888
	var n4, name, n5 = 100, "tom", 888
	fmt.Println("n4 = ",n4, "name = ",name, "n5 = ",n5)

	// 一次性声明多个变量的方式2,同样可以使用类型推导
	//输出结果为 n6 =  100 name1 =  tom~ n7 =  8888
	n6, name1, n7 := 100, "tom~", 8888
	fmt.Println("n6 = ",n6, "name1 = ",name1, "n7 = ",n7)
}