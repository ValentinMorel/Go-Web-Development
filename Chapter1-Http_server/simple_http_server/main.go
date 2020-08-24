package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	ConnHost = "localhost"
	ConnPort = "8080"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World ! ")
}

func main() {
	http.HandleFunc("/", HelloWorld)
	err := http.ListenAndServe(ConnHost+":"+ConnPort, nil)
	if err != nil {
		log.Fatal(err)
	}

}
