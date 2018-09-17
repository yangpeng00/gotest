package main

import (
	"fmt"
)

type Student struct {
	name  string
	age   int
	score map[string]int
}

func main() {
	var stmap map[int]Student
	stmap = make(map[int]Student)

	var stsc map[string]int
	stsc = make(map[string]int)
	stsc["语文"] = 1223
	stsc["数学"] = 232
	s := Student{
		name:  "www",
		age:   12,
		score: stsc}

	stmap[1] = s
	var stsc2 map[string]int
	stsc2 = make(map[string]int)
	stsc2["语文"] = 1211
	stsc2["数学"] = 2311
	s2 := Student{
		name:  "rrr",
		age:   54,
		score: stsc2}
	stmap[2] = s2

	fmt.Println(stmap)
}
