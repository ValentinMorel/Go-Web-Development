package main

import (
	"log"
	"net"
)

// put static parameters in cache
const (
	ConnType = "tcp"
	ConnHost = "localhost"
	ConnPort = "8080"
)

func main() {

	listener, err := net.Listen(ConnType, ConnHost+":"+ConnPort)
	if err != nil {
		// shutdown program
		log.Fatal(err)
	}
	// close the tcp socket when program exits
	defer listener.Close()
	// log the active connection
	log.Printf("Started listening on address : %s and port %s", ConnHost, ConnPort)

	// infinite loop, listening for incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(conn)
	}

}
