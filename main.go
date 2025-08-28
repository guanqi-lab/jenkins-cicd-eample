package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! v1.0.6")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server starting on port 28080...")
	http.ListenAndServe(":28080", nil)
}
