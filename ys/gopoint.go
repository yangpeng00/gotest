package main

import "fmt"

/*
1.指针的创建使用，指针参数等
*/
func main() {
	var a = 20
	var ia *int

	ia = &a
	fmt.Println("%x\n", &a)
	fmt.Println("%x\n", &ia)
	fmt.Println("%x\n", *ia)
	fmt.Println("_________________________________________")
	var a2 int = 20 /* 声明实际变量 */
	var ip *int     /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a2)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)

}
