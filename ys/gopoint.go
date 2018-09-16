package main

import "fmt"

/*
1.指针的创建使用，指针参数等
*/
func main() {
	var a int = 20
	var ia *int
	ia = &a
	fmt.Print(&a)
	fmt.Print("%x\n", ia)
	fmt.Print("%x\n", *ia)
}
