package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	const port string = ":8000"
	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "up and running...")
	})
	router.HandleFunc("/posts", GetPosts).Methods("GET")
	router.HandleFunc("/post", AddPost).Methods("POST")
	log.Println("server listening on port", port)
	log.Fatalln(
		http.ListenAndServe(port, router),
	)
}
