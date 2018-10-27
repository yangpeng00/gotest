package main

import (
	"fmt"
	"log"
	"net/http"
)

func init() {
	fmt.Printf("asdfad")
}
func main() {
	fmt.Println("123123")
	sss, err := http.Get("http://www.baidu.com")
	if err != nil {
		log.Println(err)
		return
	}
	defer sss.Body.Close()

	head := sss.Header

	for a, b := range head {
		fmt.Printf("a=%v,b=%v\n", a, b)
	}

}
