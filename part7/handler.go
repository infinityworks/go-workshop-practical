package practical

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sashabaranov/go-openai"
)

type request struct {
	Transcript string `json:"transcript"`
}

type response struct {
	Summary string `json:"summary"`
}

type TranscriptSummaryHandler struct {
	Summariser OpenAISummariser
}

// We can create a constructor for our http handler
func NewTranscriptSummaryHandler(token string) TranscriptSummaryHandler {
	sum := NewSummariser(token)
	return TranscriptSummaryHandler{
		Summariser: sum,
	}
}

func (h TranscriptSummaryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var requestMessage request

	err := json.NewDecoder(r.Body).Decode(&requestMessage)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// Call our new function
	summary, err := h.Summariser.Summarise(requestMessage.Transcript)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	responseMessage := response{Summary: summary}

	err = json.NewEncoder(w).Encode(responseMessage)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
}

// Create a wrapper for the openai client.
type OpenAISummariser struct {
	Client *openai.Client
}

// Create a constructor for the wrapper, in go constructors are just functions that return a type.
func NewSummariser(token string) OpenAISummariser {
	return OpenAISummariser{
		Client: openai.NewClient(token),
	}
}

// Create a summarise method
func (s OpenAISummariser) Summarise(transcript string) (string, error) {
	// Construct our prompt
	prompt := fmt.Sprintf("Write a synopsis for the following transcript:\n%s", transcript)
	// Call the openai API
	resp, err := s.Client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo16K,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	})
	// errors can be wrapped to add context.
	if err != nil {
		return "", fmt.Errorf("Call to openai failed: %w", err)
	}
	return resp.Choices[0].Message.Content, nil
}
