package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.56.1:4242")
	if err != nil {
		log.Fatal(err)
	}
	conn.Write([]byte("Privet mir"))

	data := make([]byte, 16)
	_, err = conn.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(data))
	defer conn.Close()
}
