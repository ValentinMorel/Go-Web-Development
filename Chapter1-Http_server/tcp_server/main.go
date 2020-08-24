// When the program is running, we can perform a request with the command :
// echo "message " | nc address port
// the message will appear on the server side if it's received.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// put static parameters in cache
const (
	ConnType = "tcp"
	ConnHost = "localhost"
	ConnPort = "8080"
)

func handleRequest(conn net.Conn) {
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Println("error reading : ", err.Error())
	}
	fmt.Println("Message received from the client : ", string(message))
	conn.Close()
}

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
		go handleRequest(conn)
	}

}
