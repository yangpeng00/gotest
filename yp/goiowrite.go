package main

import (
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//func main() {
//	d1 := []byte("hello\ngo\n")
//	err := ioutil.WriteFile("D:\\test\\ww.txt", d1, 0644)
//	check(err)
//}

func main() {
	b, err := ioutil.ReadFile("D:\\test\\ww.txt")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(b)
	str := string(b)
	fmt.Println(str)
}
