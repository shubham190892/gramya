package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	server()
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Default, %q", html.EscapeString(r.URL.Path))
}

func server() {
	port := ":8383"
	http.HandleFunc("/", defaultHandler)
	fmt.Printf("Http server listening at %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
