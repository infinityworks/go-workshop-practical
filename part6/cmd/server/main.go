package main

import (
	"fmt"
	practical "github.com/infinityworks/go-workshop-practical"
	"net/http"
	"os"
)

func main() {
	// Create a go router
	mux := http.NewServeMux()

	// Route all requests to our handler
	mux.Handle("/", practical.TranscriptSummaryHandler{})

	// Start the server
	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
