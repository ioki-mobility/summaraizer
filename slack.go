package summaraizer

import (
	"encoding/json"
	"io"
	"net/http"
)

// Slack is a source that fetches comments from a Slack thread.
type Slack struct {
	Token   string // The OAuth token.
	Channel string // The channel ID.
	TS      string // The timestamp of the thread.
}

// Fetch fetches comments from a Slack thread.
func (s *Slack) Fetch(writer io.Writer) error {
	return fetchAndEncode(writer, func() (Comments, error) {
		url := "https://slack.com/api/conversations.replies?channel=" + s.Channel + "&ts=" + s.TS
		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		request.Header.Set("Authorization", "Bearer "+s.Token)
		request.Header.Set("Content-Type", "application/json; charset=utf-8")

		resp, err := http.DefaultClient.Do(request)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var response slackThreadResponse
		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, err
		}

		var comments Comments
		for _, message := range response.Messages {
			comments = append(comments, Comment{
				Author: message.User,
				Body:   message.Text,
			})
		}

		return comments, nil
	})
}

type slackThreadResponse struct {
	Messages []struct {
		User string `json:"user"`
		Text string `json:"text"`
	} `json:"messages"`
}
