package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Get the port from the environment variable, default to 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	msg := os.Getenv("MESSAGE")
	if msg == "" {
		msg = "ok"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, msg)
	})

	// Start the server on the specified port
	fmt.Printf("Server is running on http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
