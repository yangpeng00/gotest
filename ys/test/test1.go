package main

import "fmt"

func main() {
	var i = 3
	var a *int
	a = &i
	as(*a)
	fmt.Print(a)
}

func as(c int) int {
	return c + 1
}
