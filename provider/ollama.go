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
	var commentsPrompt string
	for _, comment := range comments {
		commentsPrompt += fmt.Sprintf("<comment>%s</comment>", comment)
	}
	reqJson := make(map[string]any)
	reqJson["model"] = o.Model
	reqJson["prompt"] = "I give you a discussion and you give me a summary. Each comment of the discussion is wrapped in a <comment> tag. Your summary should not be longer than 1200 chars. Here is the discussion: " + commentsPrompt
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
