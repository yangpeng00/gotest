package main

import "fmt"

/*
1.go的基本类型的用法
2.常量，变量的创建使用
*/
func main() {

	var a = 31
	var aa, b, c, d int = 1, 2, 3, 4
	var (
		bb int
		cc int
		dd bool
	)
	fmt.Print(a, aa, b, c, d, bb, cc, dd)
}
