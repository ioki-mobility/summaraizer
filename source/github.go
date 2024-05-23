package source

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ioki-mobility/summaraizer"
)

// GitHub is a source that fetches comments from a GitHub issue or pull request.
type GitHub struct {
	Token       string
	RepoOwner   string
	RepoName    string
	IssueNumber string
}

// Fetch fetches comments from a GitHub issue.
func (gh *GitHub) Fetch(writer io.Writer) error {
	comments, err := gh.fetchInternal()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(writer)
	return encoder.Encode(comments)
}

func (gh *GitHub) fetchInternal() (summaraizer.Comments, error) {
	var comments summaraizer.Comments

	issue := issue{
		repo: repo{
			Owner: gh.RepoOwner,
			Name:  gh.RepoName,
		},
		Number: gh.IssueNumber,
	}

	issueResponse := fetchIssue(gh.Token, issue)

	comments = append(comments, summaraizer.Comment{
		Author: issueResponse.User.Login,
		Body:   issueResponse.Body,
	})

	if issueResponse.Comments > 0 {
		issueComments := fetchIssueComments(gh.Token, issue)
		for _, issueComments := range *issueComments {
			comments = append(comments, summaraizer.Comment{
				Author: issueComments.User.Login,
				Body:   issueComments.Body,
			})
		}
	}

	return comments, nil
}

func fetchIssue(token string, issue issue) *issueResponse {
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

	var issueResponse issueResponse
	err = json.Unmarshal(body, &issueResponse)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return &issueResponse
}

func fetchIssueComments(token string, issue issue) *issueCommentsResponse {
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

	var commentsResponse issueCommentsResponse
	err = json.Unmarshal(body, &commentsResponse)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return &commentsResponse
}

func newRequest(method, url string, token string) http.Request {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}
	addHeaders(req, token)
	return *req
}

func addHeaders(req *http.Request, token string) {
	req.Header.Add("Accept", "application/vnd.github+json")
	// Public repos doesn't need Authorization header
	if token != "" {
		req.Header.Add("Authorization", "Bearer "+token)
	}
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")
}

type repo struct {
	Owner string
	Name  string
}

type issue struct {
	repo
	Number string
}

type issueResponse struct {
	Body     string       `json:"body"`
	User     userResponse `json:"user"`
	Comments int          `json:"comments"`
}

type userResponse struct {
	Login string `json:"login"`
}

type issueCommentsResponse = []issueCommentResponse

type issueCommentResponse struct {
	Body string       `json:"body"`
	User userResponse `json:"user"`
}
