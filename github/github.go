package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchIssue(token string, issue Issue) *IssueResponse {
	issueURL := fmt.Sprintf(
		"https://api.github.com/repos/%s/%s/issues/%s",
		issue.Owner,
		issue.Name,
		issue.Number,
	)
	request := newRequest("GET", issueURL, token)
	resp, err := http.DefaultClient.Do(&request)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	var issueResponse IssueResponse
	err = json.Unmarshal(body, &issueResponse)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return &issueResponse
}

func FetchIssueComments(token string, issue Issue) *IssueCommentsResponse {
	commentsURL := fmt.Sprintf(
		"https://api.github.com/repos/%s/%s/issues/%s/comments",
		issue.Owner,
		issue.Name,
		issue.Number,
	)
	request := newRequest("GET", commentsURL, token)
	resp, err := http.DefaultClient.Do(&request)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	var commentsResponse IssueCommentsResponse
	err = json.Unmarshal(body, &commentsResponse)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return &commentsResponse
}
