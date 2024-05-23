package summaraizer

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// OpenAi is a provider that uses OpenAI API as an AI provider.
type OpenAi struct {
	Common
	ApiToken string
}

func (o *OpenAi) Summarize(reader io.Reader) (string, error) {
	return decodeAndSummarize(reader, func(comments Comments) (string, error) {
		prompt, err := resolvePrompt(o.Common.Prompt, comments)
		if err != nil {
			return "", err
		}

		request := openAiRequest{
			Model: o.Model,
			Messages: []openAiRequestMessage{
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
			"https://api.openai.com/v1/chat/completions",
			bytes.NewBuffer(reqBodyBytes),
		)
		if err != nil {
			return "", err
		}
		req.Header.Set("Authorization", "Bearer "+o.ApiToken)
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

		var response openAiResponse
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return "", err
		}

		return response.Choices[0].Message.Content, nil
	})
}

type openAiRequest struct {
	Model    string                 `json:"model"`
	Messages []openAiRequestMessage `json:"messages"`
}

type openAiRequestMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type openAiResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
			Role    string `json:"role"`
		} `json:"message"`
	} `json:"choices"`
}
