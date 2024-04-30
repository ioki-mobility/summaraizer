package summairizer

import (
	"github.com/ioki-mobility/summaraizer/github"
	"github.com/ioki-mobility/summaraizer/types"
)

type GitHubInput struct {
	Token       string
	RepoOwner   string
	RepoName    string
	IssueNumber string
}

type AiProvider interface {
	Summarize(comments types.Comments) (string, error)
}

func Summarize(gitHubInput GitHubInput, aiProvider AiProvider) (string, error) {
	comments, err := fetchDiscussion(gitHubInput)
	if err != nil {
		return "", err
	}

	summarization, err := aiProvider.Summarize(comments)
	if err != nil {
		return "", err
	}

	return summarization, nil
}

// fetchDiscussion fetches the discussion from a GitHub issue
func fetchDiscussion(
	input GitHubInput,
) (types.Comments, error) {
	var comments types.Comments

	issue := github.Issue{
		Repo: github.Repo{
			Owner: input.RepoOwner,
			Name:  input.RepoName,
		},
		Number: input.IssueNumber,
	}

	issueResponse := github.FetchIssue(input.Token, issue)

	comments = append(comments, types.Comment{
		Author: issueResponse.User.Login,
		Body:   issueResponse.Body,
	})

	if issueResponse.Comments > 0 {
		issueComments := github.FetchIssueComments(input.Token, issue)
		for _, issueComments := range *issueComments {
			comments = append(comments, types.Comment{
				Author: issueComments.User.Login,
				Body:   issueComments.Body,
			})
		}
	}

	return comments, nil
}
