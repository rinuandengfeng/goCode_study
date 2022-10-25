package main

import "fmt"

func main() {

	//写入defer关键字
	defer fmt.Println("main end1")
	defer fmt.Println("main end2") //end2先执行

	fmt.Println("main::hello go 1")
	fmt.Println("main::hello go 2")
}
