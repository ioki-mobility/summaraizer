package summaraizer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Google is a provider that uses Google as an AI provider.
type Google struct {
	Model    string // The Ai model to use.
	Prompt   string // The prompt to use for the AI model.
	ApiToken string // The API Token for Google.
}

func (a *Google) Summarize(reader io.Reader) (string, error) {
	return decodeAndSummarize(reader, func(comments Comments) (string, error) {
		prompt, err := resolvePrompt(a.Prompt, comments)
		if err != nil {
			return "", err
		}

		request := googleRequest{
			GoogleContentsRequest: []googleContentsRequest{
				{
					Parts: []googlePartsRequest{
						{
							Text: prompt,
						},
					},
				},
			},
		}
		reqBodyBytes, err := json.Marshal(request)
		if err != nil {
			return "", err
		}

		url := fmt.Sprintf(
			"https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent?key=%s",
			a.Model,
			a.ApiToken,
		)
		req, err := http.NewRequest(
			"POST",
			url,
			bytes.NewBuffer(reqBodyBytes),
		)
		if err != nil {
			return "", err
		}
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
		var response googleResponse
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return "", err
		}

		return response.Candidates[0].Content.Parts[0].Text, nil
	})
}

type googleRequest struct {
	GoogleContentsRequest []googleContentsRequest `json:"contents"`
}

type googleContentsRequest struct {
	Parts []googlePartsRequest `json:"parts"`
}

type googlePartsRequest struct {
	Text string `json:"text"`
}

type googleResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}
