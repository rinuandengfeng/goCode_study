package main

import (
	_ "E:/gocode/goStudy/5-init/lib1" //前面加_起一个匿名，不会报错
	lib2 "./E:/gocodegoStudy/5-init/lib2"  //还可以加 . 表示把这个包导入到当前包里
	
)

func main() {
	lib1.Lib1Test()
	lib2.Lib2Test()
}
