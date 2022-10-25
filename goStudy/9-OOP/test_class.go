package main

import "fmt"

//定义一个类方法，如果类名首字母大写，表示其他包也能够访问
type Hero struct {
	//如果说类的属性首字母大写，表示该属性是对外能够访问的，否则的话只能够类的内部访问
	Name  string
	Ad    int
	Level int
}

// func (this Hero) Show() {
// 	fmt.Println("Name = ", this.Name)
// 	fmt.Println("Ad = ", this.Ad)
// 	fmt.Println("Level = ", this.Level)
// }

// func (this Hero) GetName() string {
// 	fmt.Println("name = ", this.Name)
// }

// func (this Hero) NewName(newName string) {
// 	this.Name = newName
// }

//这里的Hero方法写成*Hero，不然是拷贝的副本，改变不了内部的值
func (this *Hero) Show() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("Level = ", this.Level)
}

func (this *Hero) GetName() string {
	fmt.Println("name = ", this.Name)
	return this.Name
}

func (this *Hero) SetName(newName string) {
	//this 是调用方法的对象的一个副本 (拷贝)
	this.Name = newName
}
func main() {

	//创建一个对象
	hero := Hero{Name: "shang3", Ad: 100, Level: 1}

	hero.Show()

	hero.SetName("li4")

	hero.Show()
}
