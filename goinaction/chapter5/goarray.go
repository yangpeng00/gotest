package main

import "fmt"

func main() {
	// 创建一个整型切片
	// 其长度和容量都是 5 个元素
	slice := []int{10, 20, 30, 40, 50}
	// 创建一个新切片
	// 其长度是 2 个元素，容量是 4 个元素
	newSlice := slice[1:3:3]
	// 修改 newSlice 索引为 1 的元素
	// 同时也修改了原来的 slice 的索引为 2 的元素
	newSlice[1] = 35

	//slice = append(slice, 44)
	newSlice = append(newSlice, 23)

	fmt.Println(slice)
	fmt.Println(newSlice)
}
