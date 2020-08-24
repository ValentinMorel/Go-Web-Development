package main

import (
	"crypto/subtle"
	"fmt"
	"log"
	"net/http"
)

const (
	ConnHost      = "localhost"
	ConnPort      = "8080"
	AdminUser     = "admin"
	AdminPassword = "pass"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World !")
}

func BasicAuth(handler http.HandlerFunc, realm string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || subtle.ConstantTimeCompare([]byte(user),
			[]byte(AdminUser)) != 1 || subtle.ConstantTimeCompare([]byte(pass),
			[]byte(AdminPassword)) != 1 {
			w.Header().Set("www-Authenticate", `Basic realm="`+realm+`"`)
			w.WriteHeader(401)
			w.Write([]byte("Unauthorized access to app.\n"))
			return
		}
		handler(w, r)
	}
}

func main() {
	http.HandleFunc("/", BasicAuth(HelloWorld, "Please Enter User and Pass"))
	err := http.ListenAndServe(ConnHost+":"+ConnPort, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

}
