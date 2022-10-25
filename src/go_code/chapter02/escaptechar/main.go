package main

import "fmt"  //fmt包中提供了格式化，输出，输入的函数

func main(){
	//演示转义字符的使用

	fmt.Println(" hello,world!")
	fmt.Println(" \thello,world!")  // '\t' 表示制表符
	fmt.Println("D:\\software\\Go\\api")   // '\\'表示，前面的'\'表示转义,后面的'\'表示\
	//从当前行的最前面开始输出，覆盖掉以前输出的内容
	fmt.Println("天龙八部雪山飞狐\r张飞")        // '\r' 表示，回车   输出内容“张飞龙八部雪山飞狐”  
	fmt.Println(" hello,\nworld!")      // '\n' 表示 ，换行符

	
}