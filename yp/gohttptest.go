package main

import (
	"io"
	"log"
	"net/http"
)

//方式一
//func main() {
//	http.HandleFunc("/hello",hello)
//	err := http.ListenAndServe(":8080",nil)
//	if err != nil {
//		log.Panic(err)
//	}
//}
//func hello(w http.ResponseWriter,r *http.Request)  {
//	io.WriteString(w,"hello")
//}

//方式二
func main() {
	http.Handle("/hello", &ServeMux{})
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Panic(err)
	}
}

type ServeMux struct {
}

func (p *ServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world")
}
