package main

import (
	"GeekBrains-1/lessons/8/001/mycontext"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", mycontext.Handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
