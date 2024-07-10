package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/ioki-mobility/summaraizer"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "summaraizer",
	Short: "Summarizes comments",
	Long:  "A tool to summarize comments from various sources using AI.",
}

func NewRootCmd() *cobra.Command {
	rootCmd.AddCommand(githubCmd(), redditCmd(), gitlabCmd(), slackCmd())
	rootCmd.AddCommand(ollamaCmd(), openaiCmd(), anthropicCmd())
	return rootCmd
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

func slackCmd() *cobra.Command {
	flagToken := "token"
	flagChannel := "channel"
	flagTs := "ts"
	var cmd = &cobra.Command{
		Use:   "slack",
		Short: "Summarizes using Slack as source",
		Long:  "Summarizes using Slack as source",
		RunE: func(cmd *cobra.Command, args []string) error {
			token, _ := cmd.Flags().GetString(flagToken)
			channel, _ := cmd.Flags().GetString(flagChannel)
			ts, _ := cmd.Flags().GetString(flagTs)

			s := &summaraizer.Slack{
				Token:   token,
				Channel: channel,
				TS:      ts,
			}
			return fetch(s)
		},
	}
	cmd.Flags().String(flagToken, "", "The Slack token.")
	cmd.MarkFlagRequired(flagToken)
	cmd.Flags().String(flagChannel, "", "The channel ID of the Slack thread.")
	cmd.MarkFlagRequired(flagChannel)
	cmd.Flags().String(flagTs, "", "The timestamp of the Slack thread.")
	cmd.MarkFlagRequired(flagTs)

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

func anthropicCmd() *cobra.Command {
	flagToken := "token"
	cmd := &cobra.Command{
		Use:   "anthropic",
		Short: "Summarizes using Anthropic AI",
		Long:  "Summarizes using Anthropic AI.",
		RunE: func(cmd *cobra.Command, args []string) error {
			aiModel, _ := cmd.Flags().GetString(aiFlagModel)
			aiPrompt, _ := cmd.Flags().GetString(aiFlagPrompt)
			apiToken, _ := cmd.Flags().GetString(flagToken)

			p := &summaraizer.Anthropic{
				Model:    aiModel,
				Prompt:   aiPrompt,
				ApiToken: apiToken,
			}

			return summarize(p)
		},
	}
	cmd.Flags().String(aiFlagModel, "claude-3-haiku-20240307", "The AI model to use")
	cmd.Flags().String(aiFlagPrompt, defaultPromptTemplate, "The prompt to use for the AI model")
	cmd.Flags().String(flagToken, "", "The API Token for Anthropic")
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
