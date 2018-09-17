package main

import (
	"log"
	"net/http"
	"time"
)

type TimeHandlerw struct {
	format string
}

func (th *TimeHandlerw) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(th.format)
	w.Write([]byte("The time is: " + tm))
}

func main() {
	mux := http.NewServeMux()

	th := &TimeHandlerw{format: time.RFC1123}
	mux.Handle("/time", th)

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}
