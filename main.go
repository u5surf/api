package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// handler := http.NewServeMux()              // router to expose api
	handler := mux.NewRouter()
	handler.HandleFunc("/api/hello", SayHello) // hello world endpoint
	fmt.Println("Server starting on localhost:8080...")
	http.ListenAndServe(":8080", handler) // listen to requests on localhost:8080
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!\n")
}
