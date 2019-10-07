package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type ScreenShot struct {
	Width    int64  `json:"width"`
	Height   int64  `json:"height"`
	Filename string `json:"filename"`
}

type Ticket struct {
	ID         int64      `json:"id"`
	Name       string     `json:"name"`
	URL        string     `json:"url"`
	Screenshot ScreenShot `json:"screenshot"`
}

func main() {
	// handler := http.NewServeMux()              // router to expose api
	handler := mux.NewRouter()
	handler.HandleFunc("/api/hello", SayHello) // hello world endpoint
	handler.HandleFunc("/api/createClient", CreateClient)
	fmt.Println("Server starting on localhost:8080...")
	http.ListenAndServe(":8080", handler) // listen to requests on localhost:8080
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!\n")
}

func CreateClient(respond http.ResponseWriter, request *http.Request) {
	var ticket Ticket
	json.NewDecoder(request.Body).Decode(&ticket)
	fmt.Println(ticket)
	fmt.Fprintf(respond, "We will create the honey-client instance here.\n")
}
