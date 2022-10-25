package main

import "fmt"

//const 来定义枚举类型
const (
	//可以在const()添加一个关键字 iota ，每行iota都会累加1，第一行的iota的默认值是0
	BEIJING = 10 * iota //iota =0
	SHANGHAI
	SHENZHEN
)

const (
	a, b = iota + 1, iota + 2 //iota =0   a = iota +1, b = iota +2
	c, d                      //iota =1   c = iota +1, d = iota +2
	e, f                      //iota=2    e = iota +1, f = iota +2
	g, h = iota * 2, iota * 3 //iota =3   g = iota * 2, h = iota * 3
	i, k                      //iota =4   i = iota * 2, k = iota * 3
)

func main() {
	//常量（只读属性）
	const lenght int = 10
	fmt.Println("lenght = ", lenght)

	//lenght = 100  //常量不允许修改的。

	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("SHENZHEN = ", SHENZHEN)
	fmt.Println("a = ", a, "b = ", b)
	fmt.Println("c = ", c, "d = ", d)
	fmt.Println("e = ", e, "f =", f)
	fmt.Println("g = ", g, "h= ", h)
	fmt.Println("i = ", i, "k = ", k)

	// iota只能出现在const中，iota 只能在const中使用

}
