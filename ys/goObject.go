package main

import "fmt"

/*
对象的创建使用
*/
type person struct {
	name string
	age  int
	cj   map[string]int
}

func main() {
	var person1 person
	var person2 person

	person1.name = "李四"
	person1.age = 31
	person1.cj = make(map[string]int)
	person1.cj["数学"] = 32
	person1.cj["语文"] = 95

	person2.name = "张三"
	person2.age = 23
	person2.cj = make(map[string]int)
	person2.cj["数学"] = 99
	person2.cj["语文"] = 88

	var per map[int]person
	per = make(map[int]person)
	per[1] = person1
	per[2] = person2
	fmt.Print(per)
}
