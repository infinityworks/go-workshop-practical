package practical

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Create a request object
type request struct {
	Transcript string `json:"transcript"`
}

// Create a response object
type response struct {
	Summary string `json:"summary"`
}

type TranscriptSummaryHandler struct{}

func (h TranscriptSummaryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// We are requiring a request body so let's only allow POST
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// Create a variable for our request body
	var requestMessage request

	// Decode the body using a json decoder (we pass a pointer as the decode function modifies it)
	err := json.NewDecoder(r.Body).Decode(&requestMessage)
	// Return a 500 if decoding fails
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println("Transcript received:", requestMessage.Transcript)
	responseMessage := response{Summary: "Placeholder"}

	// Encode the json response, it can be written straight to the ResponseWriter
	err = json.NewEncoder(w).Encode(responseMessage)
	// Return a 500 if encoding fails
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
