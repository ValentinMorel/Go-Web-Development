package main

import (
	"log"
	"net"
)

const (
	ConnType = "tcp"
	ConnHost = "localhost"
	ConnPort = "8080"
)

func main() {

	listener, err := net.Listen(ConnType, ConnHost+":"+ConnPort)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	log.Printf("Started listening on address : %s and port %s", ConnHost, ConnPort)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(conn)
	}

}
