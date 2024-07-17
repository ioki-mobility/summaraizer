package summaraizer

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// Anthropic is a provider that uses Anthropic as an AI provider.
type Anthropic struct {
	Model    string // The Ai model to use.
	Prompt   string // The prompt to use for the AI model.
	ApiToken string // The API Token for Anthropic.
}

func (a *Anthropic) Summarize(reader io.Reader) (string, error) {
	return decodeAndSummarize(reader, func(comments Comments) (string, error) {
		prompt, err := resolvePrompt(a.Prompt, comments)
		if err != nil {
			return "", err
		}

		request := anthropicRequest{
			Model:     a.Model,
			MaxTokens: 4096,
			Messages: []anthropicMessageRequest{
				{
					Role:    "user",
					Content: prompt,
				},
			},
		}
		reqBodyBytes, err := json.Marshal(request)
		if err != nil {
			return "", err
		}

		req, err := http.NewRequest(
			"POST",
			"https://api.anthropic.com/v1/messages",
			bytes.NewBuffer(reqBodyBytes),
		)
		if err != nil {
			return "", err
		}
		req.Header.Set("x-api-key", a.ApiToken)
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		var response anthropicResponse
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return "", err
		}

		return response.Content[0].Text, nil
	})
}

type anthropicRequest struct {
	Model     string                    `json:"model"`
	MaxTokens int                       `json:"max_tokens"`
	Messages  []anthropicMessageRequest `json:"messages"`
}

type anthropicMessageRequest struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type anthropicResponse struct {
	Content []struct {
		Text string `json:"text"`
	} `json:"content"`
}
