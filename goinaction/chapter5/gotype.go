package main

import "fmt"

func main() {

	fmt.Println(isShellSpecialVar(7))
}

func isShellSpecialVar(c uint8) bool {
	switch c {
	case '*', '#', '$', '@', '!', '?', '0', '1', '2', '3', '4', '5',
		'6', '7', '8', '9':
		return true
	}
	return false
}
