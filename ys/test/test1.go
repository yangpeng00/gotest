package main

import "fmt"

func main() {
	var i = 3
	var a *int
	a = &i
	b := as(*a)
	fmt.Print(b)
}

func as(c int) int {
	return c + 1
}
