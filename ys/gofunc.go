package main

import "fmt"

/*
1.创建函数    创建有参数有返回值的函数，，参数有各种类型，返回参数有对个或各种类型
*/
func lll(a string, b int) (string, int) {
	//	fmt.Printf("%s的年龄是%d",a,b)
	a = "asedf"
	b = 1234
	return a, b
}
func main() {
	c := "张华"
	d := 312
	c, d = lll(c, d)
	fmt.Println(c, d)
}
