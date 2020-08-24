package main

import (
	"io"
	"net/http"

	"github.com/gorilla/handlers"
)

const (
	ConnHost = "localhost"
	ConnPort = "8080"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World !")
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloWorld)

	err := http.ListenAndServe(ConnHost+":"+ConnPort, handlers.CompressHandler(mux))
	if err != nil {
		panic(err)
	}
}
