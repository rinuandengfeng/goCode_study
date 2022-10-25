package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	var s,sep string
	for i := 1; i < len(os.Args); i++{
		s += sep + os.Args[i]
		sep = ""
	}
	fmt.Println(s)
	fmt.Println(len(os.Args))

	//打印os.Args[0]
	fmt.Println(os.Args[0])

	//strings.Join和直接Println的输出结果区别
	fmt.Println(os.Args[1:])
	fmt.Println("1")
	fmt.Println(strings.Join(os.Args[1:],""))

	//打印value和index，每个value和index显示在一行
	for index, value := range os.Args{
		fmt.Println("index=", index )
		fmt.Println("value=", value )
	}
	
	
}