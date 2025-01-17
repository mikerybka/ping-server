package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ok")
}

func main() {
	http.HandleFunc("/", handler)

	// Get the port from the environment variable, default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Start the server on the specified port
	fmt.Printf("Server is running on http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
