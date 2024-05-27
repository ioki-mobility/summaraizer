package summaraizer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Ollama is a provider that uses Ollama as an AI provider.
type Ollama struct {
	Model  string // The Ai model to use.
	Prompt string // The prompt to use for the AI model.
	Url    string // The URL where Ollama is accessible.
}

func (o *Ollama) Summarize(reader io.Reader) (string, error) {
	return decodeAndSummarize(reader, func(comments Comments) (string, error) {
		prompt, err := resolvePrompt(o.Prompt, comments)
		if err != nil {
			return "", err
		}

		request := ollamaRequest{
			Model:  o.Model,
			Prompt: prompt,
			Stream: false,
		}
		reqBodyBytes, err := json.Marshal(request)
		if err != nil {
			return "", err
		}

		resp, err := http.Post(
			fmt.Sprintf("%s/api/generate", o.Url),
			"application/json",
			bytes.NewBuffer(reqBodyBytes),
		)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}

		var response ollamaResponse
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return "", err
		}

		return response.Response, nil
	})
}

type ollamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type ollamaResponse struct {
	Response string `json:"response"`
}
