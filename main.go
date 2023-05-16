package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting")
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong!")
}
