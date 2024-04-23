package main

import (
	"flag"
	"fmt"
	"os"

	summaraizer "github.com/ioki-mobility/summaraizer"
)

func main() {
	aiProviderName := flag.String("ai-provider", "ollama", "AI Provider Name")
	aiModel := flag.String("ai-model", "mistral:7b", "AI Model")
	token := flag.String("token", "", "GitHub Token (can be empty for public repositories)")
	owner := flag.String("owner", "", "GitHub Repository Owner")
	repo := flag.String("repo", "", "GitHub Repository Name")
	issueNumber := flag.String("issue-number", "", "GitHub Issue Number")

	flag.Parse()

	if *owner == "" || *repo == "" || *issueNumber == "" {
		fmt.Println("Usage: summaraizer -owner <owner> -repo <repo> -issue-number <issueNumber> [-token <token>]")
		os.Exit(1)
	}

	input := summaraizer.Input{
		AiInput: summaraizer.AiInput{
			AiProviderName: summaraizer.AiProviderName(*aiProviderName),
			AiModel:        *aiModel,
		},
		GitHubInput: summaraizer.GitHubInput{
			Token:       *token,
			RepoOwner:   *owner,
			RepoName:    *repo,
			IssueNumber: *issueNumber,
		},
	}
	summarization, err := summaraizer.Summarize(input)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(summarization)
}
