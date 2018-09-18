package main

import "fmt"

func main() {
	go say()
}

func say() {
	fmt.Println("hello world")
}
