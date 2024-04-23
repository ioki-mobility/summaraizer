package summairizer

import (
	"errors"

	"github.com/ioki-mobility/summaraizer/github"
	"github.com/ioki-mobility/summaraizer/provider"
	"github.com/ioki-mobility/summaraizer/types"
)

type AiProviderName string

const (
	Ollama  AiProviderName = "ollama"
	OpenAI  AiProviderName = "openai"
	Mistral AiProviderName = "mistral"
)

type AiInput struct {
	AiProviderName AiProviderName
	AiModel        string
}

type GitHubInput struct {
	Token       string
	RepoOwner   string
	RepoName    string
	IssueNumber string
}

type Input struct {
	AiInput
	GitHubInput
}

type AiProvider interface {
	Summarize(model string, comments types.Comments) (string, error)
}

func Summarize(input Input) (string, error) {
	comments, err := fetchDiscussion(input.GitHubInput)
	if err != nil {
		return "", err
	}

	aiProvider, err := findAiProvider(input.AiProviderName)
	if err != nil {
		return "", err
	}

	summarization, err := aiProvider.Summarize(input.AiModel, comments)
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

func findAiProvider(model AiProviderName) (AiProvider, error) {
	switch model {
	case OpenAI:
		return provider.OpenAi{}, nil
	case Mistral:
		return provider.Mistral{}, nil
	case Ollama:
		return provider.Ollama{}, nil
	}
	return nil, errors.New("invalid AiProvider: " + string(model))
}
