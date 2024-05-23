package summaraizer

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// GitLab is a source that fetches comments from a GitLab issue or merge request.
type GitLab struct {
	Token       string
	RepoOwner   string
	RepoName    string
	IssueNumber string
	Url         string // The (base) URL of the GitLab instance.
}

// Fetch fetches comments from a GitLab issue.
func (gl *GitLab) Fetch(writer io.Writer) error {
	return fetchAndEncode(writer, func() (Comments, error) {
		var comments Comments

		issue := issue{
			repo: repo{
				Owner: gl.RepoOwner,
				Name:  gl.RepoName,
			},
			Number: gl.IssueNumber,
		}

		issueResponse, err := fetchGitLabIssue(gl, issue)
		if err != nil {
			return nil, err
		}

		comments = append(comments, Comment{
			Author: issueResponse.Author.Username,
			Body:   issueResponse.Description,
		})

		issueComments, err := fetchGitLabIssueNotes(gl, issue)
		if err != nil {
			return nil, err
		}
		for _, issueComments := range *issueComments {
			comments = append(comments, Comment{
				Author: issueComments.Author.Username,
				Body:   issueComments.Body,
			})
		}

		return comments, nil
	})
}

func fetchGitLabIssue(gl *GitLab, issue issue) (*gitLabIssueResponse, error) {
	issueURL := fmt.Sprintf(
		"%s/api/v4/projects/%s/issues/%s",
		gl.Url,
		url.QueryEscape(fmt.Sprintf("%s/%s", issue.Owner, issue.Name)),
		issue.Number,
	)
	request := newGitLabRequest("GET", issueURL, gl.Token)
	resp, err := http.DefaultClient.Do(&request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var issueResponse gitLabIssueResponse
	err = json.Unmarshal(body, &issueResponse)
	if err != nil {
		return nil, err
	}

	return &issueResponse, nil
}

func fetchGitLabIssueNotes(gl *GitLab, issue issue) (*gitLabIssueNotesResponse, error) {
	commentsURL := fmt.Sprintf(
		"%s/api/v4/projects/%s/issues/%s/notes?per_page=100",
		gl.Url,
		url.QueryEscape(fmt.Sprintf("%s/%s", issue.Owner, issue.Name)),
		issue.Number,
	)
	request := newGitLabRequest("GET", commentsURL, gl.Token)
	resp, err := http.DefaultClient.Do(&request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var commentsResponse gitLabIssueNotesResponse
	err = json.Unmarshal(body, &commentsResponse)
	if err != nil {
		return nil, err
	}

	return &commentsResponse, nil
}

func newGitLabRequest(method, url string, token string) http.Request {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}
	addGitLabHeaders(req, token)
	return *req
}

func addGitLabHeaders(req *http.Request, token string) {
	req.Header.Add("PRIVATE-TOKEN", token)
}

type gitLabIssueResponse struct {
	Description string               `json:"description"`
	Author      gitLabAuthorResponse `json:"author"`
}

type gitLabAuthorResponse struct {
	Username string `json:"username"`
}

type gitLabIssueNotesResponse = []gitLabIssueNoteResponse

type gitLabIssueNoteResponse struct {
	Body   string               `json:"body"`
	Author gitLabAuthorResponse `json:"author"`
}
