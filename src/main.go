package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing request in handleFunction\n", r.URL, r.Body)
	fmt.Fprint(w, "<h1> Hello NGINX! \nFinaly!</h1>")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		handlerFunc(w, r)
	})
	fmt.Println("Starting the server on :3002...")
	http.ListenAndServe(":3002", nil)
}
