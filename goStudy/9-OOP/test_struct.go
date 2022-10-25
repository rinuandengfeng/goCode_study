package main

import "fmt"

//声明一种行为的数据类型 myint ，是int的一个别名
type myint int

//定义一个结构体
type Book struct {
	title string
	auth  string
}

func changBook(book Book) {
	//传递一个book的副本
	book.auth = "6666"
}

func changBook2(book *Book) {
	//指针传递
	book.auth = "7777"
}

func main() {

	var a myint = 10
	fmt.Println("a = ", a)
	fmt.Printf("type of a  = %T\n", a)

	fmt.Println("----------")

	var book1 Book
	book1.title = "Golang"
	book1.auth = "xiaoliu"

	fmt.Printf("%v\n", book1)

	changBook(book1)
	fmt.Printf("%v\n", book1)

	//传递book1的地址
	changBook2(&book1)
	fmt.Printf("%v\n", book1)

}
