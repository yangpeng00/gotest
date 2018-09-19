package main

import "fmt"

func main() {
	slice := []int{10, 23, 34, 45, 65}

	for i, v := range slice {
		fmt.Printf("下标%d,值为%d\n", i, v)
	}
}
