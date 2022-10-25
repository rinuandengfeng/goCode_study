package main

import "fmt"

// 返回一个值
func foo1(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	c := 100
	return c
}

//返回多个返回值，匿名
func foo2(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	return 666, 777
}

//返回多个返回值，有形参名称的
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("-----foo3-----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	//r1, r2 属于foo3 的形参 ，初始化默认的值是0
	//r1, r2 作用域空间 是foo3 整个函数体的{}空间
	fmt.Println("r1 = ", r1) //默认等于r1 = 0 , r2 = 0
	fmt.Println("r2 = ", r2)
	//给有名称的返回值赋值
	r1 = 1000
	r2 = 2000

	return
	// return 1000,2000   这样的写法也可以

}

func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("-----foo4-----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	//给有变量的赋值
	r1 = 1000
	r2 = 2000

	return
}

func main() {
	c := foo1("abc", 555)
	fmt.Println("c = ", c)

	ret1, ret2 := foo2("hahah", 999)
	fmt.Println("ret1 = ", ret1)
	fmt.Println("ret2 = ", ret2)

	r1, r2 := foo3("foo3", 333)
	fmt.Println("r1 = ", r1, "r2 = ", r2)

	re1, re2 := foo4("foo4", 444)
	fmt.Println("re1 = ", re1, "re2 = ", re2)
}
