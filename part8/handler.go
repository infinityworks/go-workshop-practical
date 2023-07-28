package practical

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sashabaranov/go-openai"
	"golang.org/x/sync/errgroup"
)

type request struct {
	Transcripts []string `json:"transcripts"`
}

type response struct {
	Summaries []string `json:"summaries"`
}

// We can use an interface to allow for easier testing of our http handler

type Summariser interface {
	Summarise(transcript string) (summary string, err error)
	SummariseBatch(transcripts []string) (summaries []string, err error)
}

type TranscriptSummaryHandler struct {
	Summariser Summariser
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
		return
	}

	// Call our new function
	summaries, err := h.Summariser.SummariseBatch(requestMessage.Transcripts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseMessage := response{Summaries: summaries}

	err = json.NewEncoder(w).Encode(responseMessage)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
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
		Model: openai.GPT3Dot5Turbo,
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

func (s OpenAISummariser) SummariseBatch(transcripts []string) ([]string, error) {
	summaries := make([]string, len(transcripts))
	var eg errgroup.Group
	for i, t := range transcripts {
		i, t := i, t
		eg.Go(func() error {
			s, err := s.Summarise(t)
			if err != nil {
				return fmt.Errorf("Failed on transcript %d: %w", i, err)
			}
			summaries[i] = s
			return nil
		})
	}
	err := eg.Wait()
	if err != nil {
		return nil, err
	}
	return summaries, nil
}
