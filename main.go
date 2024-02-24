package main

import (
	"fmt"
	"net/http"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/a", HelloWorldHandler)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
