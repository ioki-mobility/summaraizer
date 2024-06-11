package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ioki-mobility/summaraizer"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd.AddCommand(githubCmd(), redditCmd(), gitlabCmd())
	rootCmd.AddCommand(ollamaCmd(), mistralCmd(), openaiCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "summaraizer",
	Short: "Summarizes comments",
	Long:  "A tool to summarize comments from various sources using AI.",
}

func githubCmd() *cobra.Command {
	flagIssue := "issue"
	flagToken := "token"
	var cmd = &cobra.Command{
		Use:   "github",
		Short: "Summarizes using GitHub as source",
		Long:  "Summarizes using GitHub as source",
		RunE: func(cmd *cobra.Command, args []string) error {
			issue, _ := cmd.Flags().GetString(flagIssue)
			githubIssueParts := strings.Split(issue, "/")
			token, _ := cmd.Flags().GetString(flagToken)

			s := &summaraizer.GitHub{
				Token:       token,
				RepoOwner:   githubIssueParts[0],
				RepoName:    githubIssueParts[1],
				IssueNumber: githubIssueParts[2],
			}
			return fetch(s)
		},
	}
	cmd.Flags().String(flagIssue, "", "The GitHub issue to summarize. Use the format owner/repo/issue_number.")
	cmd.MarkFlagRequired(flagIssue)
	cmd.Flags().String(flagToken, "", "The GitHub token. Only required for private repositories.")

	return cmd
}

func redditCmd() *cobra.Command {
	flagPost := "post"
	cmd := &cobra.Command{
		Use:   "reddit",
		Short: "Summarizes using Reddit as source",
		Long:  "Summarizes using Reddit as source.",
		RunE: func(cmd *cobra.Command, args []string) error {
			post, _ := cmd.Flags().GetString(flagPost)

			s := &summaraizer.Reddit{
				UrlPath: post,
			}
			return fetch(s)
		},
	}
	cmd.Flags().String(flagPost, "", "The Reddit post to summarize. Use the URL path. Everything after reddit.com.")
	cmd.MarkFlagRequired(flagPost)
	return cmd
}

func gitlabCmd() *cobra.Command {
	flagIssue := "issue"
	flagToken := "token"
	flagUrl := "url"
	var cmd = &cobra.Command{
		Use:   "gitlab",
		Short: "Summarizes using GitLab as source",
		Long:  "Summarizes using GitLab as source",
		RunE: func(cmd *cobra.Command, args []string) error {
			issue, _ := cmd.Flags().GetString(flagIssue)
			githubIssueParts := strings.Split(issue, "/")
			token, _ := cmd.Flags().GetString(flagToken)
			url, _ := cmd.Flags().GetString(flagUrl)

			s := &summaraizer.GitLab{
				Token:       token,
				RepoOwner:   githubIssueParts[0],
				RepoName:    githubIssueParts[1],
				IssueNumber: githubIssueParts[2],
				Url:         url,
			}
			return fetch(s)
		},
	}
	cmd.Flags().String(flagIssue, "", "The GitLab issue to summarize. Use the format owner/repo/issue_number.")
	cmd.MarkFlagRequired(flagIssue)
	cmd.Flags().String(flagToken, "", "The GitLab token.")
	cmd.MarkFlagRequired(flagToken)
	cmd.Flags().String(flagUrl, "https://gitlab.com", "The URL of the GitLab instance.")

	return cmd
}

const (
	aiFlagModel  string = "model"
	aiFlagPrompt string = "prompt"
)

func ollamaCmd() *cobra.Command {
	flagUrl := "url"
	var cmd = &cobra.Command{
		Use:   "ollama",
		Short: "Summarizes using Ollama AI",
		Long:  "Summarizes using Ollama AI.",
		RunE: func(cmd *cobra.Command, args []string) error {
			aiModel, _ := cmd.Flags().GetString(aiFlagModel)
			aiPrompt, _ := cmd.Flags().GetString(aiFlagPrompt)
			url, _ := cmd.Flags().GetString(flagUrl)

			p := &summaraizer.Ollama{
				Model:  aiModel,
				Prompt: aiPrompt,
				Url:    url,
			}

			return summarize(p)
		},
	}
	cmd.Flags().String(aiFlagModel, "gemma:2b", "The AI model to use")
	cmd.Flags().String(aiFlagPrompt, defaultPromptTemplate, "The prompt to use for the AI model")
	cmd.Flags().String(flagUrl, "http://localhost:11434", "The URl where ollama is accessible")
	return cmd
}

func mistralCmd() *cobra.Command {
	flagToken := "token"
	var cmd = &cobra.Command{
		Use:   "mistral",
		Short: "Summarizes using Mistral AI",
		Long:  "Summarizes using Mistral AI.",
		RunE: func(cmd *cobra.Command, args []string) error {
			aiModel, _ := cmd.Flags().GetString(aiFlagModel)
			aiPrompt, _ := cmd.Flags().GetString(aiFlagPrompt)
			apiToken, _ := cmd.Flags().GetString(flagToken)

			p := &summaraizer.Mistral{
				Model:    aiModel,
				Prompt:   aiPrompt,
				ApiToken: apiToken,
			}

			return summarize(p)
		},
	}
	cmd.Flags().String(aiFlagModel, "mistral:7b", "The AI model to use")
	cmd.Flags().String(aiFlagPrompt, defaultPromptTemplate, "The prompt to use for the AI model")
	cmd.Flags().String(flagToken, "", "The API Token for Mistral")
	cmd.MarkFlagRequired(flagToken)
	return cmd
}

func openaiCmd() *cobra.Command {
	flagToken := "token"
	cmd := &cobra.Command{
		Use:   "openai",
		Short: "Summarizes using OpenAI AI",
		Long:  "Summarizes using OpenAI AI.",
		RunE: func(cmd *cobra.Command, args []string) error {
			aiModel, _ := cmd.Flags().GetString(aiFlagModel)
			aiPrompt, _ := cmd.Flags().GetString(aiFlagPrompt)
			apiToken, _ := cmd.Flags().GetString(flagToken)

			p := &summaraizer.OpenAi{
				Model:    aiModel,
				Prompt:   aiPrompt,
				ApiToken: apiToken,
			}

			return summarize(p)
		},
	}
	cmd.Flags().String(aiFlagModel, "gpt-3.5-turbo", "The AI model to use")
	cmd.Flags().String(aiFlagPrompt, defaultPromptTemplate, "The prompt to use for the AI model")
	cmd.Flags().String(flagToken, "", "The API Token for OpenAI")
	cmd.MarkFlagRequired(flagToken)
	return cmd
}

var defaultPromptTemplate = `
I give you a discussion and you give me a summary.
Each comment of the discussion is wrapped in a <comment> tag.
Your summary should not be longer than 1200 chars.
Here is the discussion:
{{ range $comment := . }}
<comment>{{ $comment.Body }}</comment>
{{end}}
`

func fetch(s summaraizer.CommentSource) error {
	err := s.Fetch(os.Stdout)
	if err != nil {
		return err
	}
	return nil
}

func summarize(p summaraizer.Summarizer) error {
	summarization, err := p.Summarize(os.Stdin)
	if err != nil {
		return err
	}
	fmt.Println(summarization)
	return nil
}
