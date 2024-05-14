package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ioki-mobility/summaraizer"
)

// Ollama is a provider that uses Ollama as an AI provider.
type Ollama struct {
	Common
	Url string
}

func (o *Ollama) Summarize(
	comments summaraizer.Comments,
) (string, error) {
	prompt, err := resolvePrompt(o.Common.Prompt, comments)
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
}

type ollamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type ollamaResponse struct {
	Response string `json:"response"`
}
