package summaraizer

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Mistral struct {
	Model    string // The Ai model to use.
	Prompt   string // The prompt to use for the AI model.
	ApiToken string // The API Token for Mistral.
}

func (m *Mistral) Summarize(reader io.Reader) (string, error) {
	return decodeAndSummarize(reader, func(comments Comments) (string, error) {
		prompt, err := resolvePrompt(m.Prompt, comments)
		if err != nil {
			return "", err
		}

		request := mistralRequest{
			Model: m.Model,
			Messages: []mistralMessage{
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
			"https://api.mistral.ai/v1/chat/completions",
			bytes.NewBuffer(reqBodyBytes),
		)
		if err != nil {
			return "", err
		}
		req.Header.Set("Authorization", "Bearer "+m.ApiToken)
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
		var response mistralResponse
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return "", err
		}

		return response.Choices[0].Message.Content, nil
	})
}

type mistralRequest struct {
	Model    string           `json:"model"`
	Messages []mistralMessage `json:"messages"`
}

type mistralMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type mistralResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}
