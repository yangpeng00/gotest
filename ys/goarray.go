package main

import (
	"fmt"
)

/*
数组的创建使用，循环，增删改查
*/
func main() {
	var num [10]int
	var i, j int
	for i = 0; i < 10; i++ {
		num[i] = i + 10
	}
	for j = 0; j < 10; j++ {
		fmt.Printf("num[%d] = %d\n", j, num[j])
	}
	leng := len(num)
	fmt.Println(leng)
	num[3] = 5
	fmt.Println(num)
	s := num[2:5]
	fmt.Println(s)
	s = append(s, 6)
	fmt.Println(s)
	m := 1
	s = append(s[:m], s[m+1:]...)
	fmt.Println(s)
}
