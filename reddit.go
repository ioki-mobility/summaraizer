package summaraizer

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Reddit is a source that fetches comments from a Reddit post.
type Reddit struct {
	UrlPath string // The path to the Reddit post. Everything after `https://www.reddit.com`.
}

// Fetch fetches comments from a Reddit post.
func (r *Reddit) Fetch(writer io.Writer) error {
	return fetchAndEncode(writer, func() (Comments, error) {
		request, err := http.NewRequest("GET", "https://www.reddit.com/"+r.UrlPath+".json", nil)
		request.Header.Set("User-Agent", "Go:summaraizer:0.0.0")
		if err != nil {
			return nil, err
		}
		resp, err := http.DefaultClient.Do(request)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var response []redditResponse
		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, err
		}

		var comments Comments
		firstPost := response[0].Data.Children[0].Data
		fistPostBody := fmt.Sprintf("%s\n\n%s", firstPost.Title, firstPost.Selftext)
		comments = append(comments, Comment{
			Author: firstPost.Author,
			Body:   fistPostBody,
		})
		if len(response) == 2 {
			for _, child := range response[1].Data.Children {
				comments = append(comments, Comment{
					Author: child.Data.Author,
					Body:   child.Data.Body,
				})
			}
		}
		return comments, nil
	})
}

type redditResponse struct {
	Data struct {
		Children []struct {
			Data struct {
				Author   string `json:"author"`
				Title    string `json:"title,omitempty"`
				Selftext string `json:"selftext,omitempty"`
				Body     string `json:"body,omitempty"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}
