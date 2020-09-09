package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	listen, err := net.Listen("tcp", ":4242")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(listen)
	log.Println("Starting server")
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}
		data := make([]byte, 16)
		_, err = conn.Read(data)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("got data from connection", string(data))
		conn.Write(data)
		conn.Close()
	}

}
