package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	router := &http.Server{
		Addr:    ":8085",
		Handler: &RouterMux{},
	}
	err := router.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

type RouterMux struct {
}

func (p *RouterMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//可以使用map将路径与响应函数管理起来
	//map[string]func(http.ResponseWriter,*http.Request)
	switch r.URL.Path {
	case "/hello":
		io.WriteString(w, "hello router")
	case "/hello2":
		io.WriteString(w, "hello world !!!")
	default:
		io.WriteString(w, r.URL.Path)
	}
}
