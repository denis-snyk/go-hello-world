package main

import (
	"fmt"
	"net/http"
	"os"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/a", HelloWorldHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Server is running on http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}
