package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("some connected")
	})
	err := http.ListenAndServe(":4242", nil)
	if err != nil {
		log.Fatal(err)
	}
}
