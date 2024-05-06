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

	reqJson := make(map[string]any)
	reqJson["model"] = o.Model
	reqJson["prompt"] = prompt
	reqJson["stream"] = false
	reqBodyBytes, err := json.Marshal(reqJson)
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
		fmt.Println("Error:", err)
		return "", nil
	}

	var responseJson map[string]any
	err = json.Unmarshal(respBody, &responseJson)
	if err != nil {
		fmt.Println("Error:", err)
		return "", nil
	}

	return responseJson["response"].(string), nil
}
