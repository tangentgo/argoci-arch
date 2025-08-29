package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request")
	fmt.Fprintf(w, "Hello from Argo Workflows CI!")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
