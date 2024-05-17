package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ioki-mobility/summaraizer"
	"github.com/ioki-mobility/summaraizer/provider"
	"github.com/ioki-mobility/summaraizer/source"
	"github.com/spf13/cobra"
)

func main() {
	sourceCmds := []*cobra.Command{githubCmd(), redditCmd()}
	for _, cmd := range sourceCmds {
		cmd.AddCommand(ollamaCmd())
		cmd.AddCommand(mistralCmd())
		cmd.AddCommand(openaiCmd())
	}

	rootCmd.AddCommand(sourceCmds...)

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

type flag struct {
	Name        string
	Required    bool
	Description string
	Default     string
}

type gitHubFlag string

const (
	gitHubFlagIssue       gitHubFlag = "issue"
	gitHubFlagSourceToken gitHubFlag = "source-token"
)

func githubCmd() *cobra.Command {
	flags := []flag{
		{
			Name:        string(gitHubFlagIssue),
			Required:    true,
			Description: "The GitHub issue to summarize. Use the format owner/repo/issue_number.",
		},
		{
			Name:        string(gitHubFlagSourceToken),
			Required:    false,
			Description: "The GitHub token to use as source",
		},
	}
	return createDefaultCmd(
		"github",
		"Summarizes using GitHub as source",
		"Summarizes using GitHub as source",
		flags,
	)
}

type redditFlag string

const (
	redditFlagPost redditFlag = "post"
)

func redditCmd() *cobra.Command {
	flags := []flag{
		{
			Name:        string(redditFlagPost),
			Required:    true,
			Description: "The Reddit post to summarize. Use the URL path. Everything after reddit.com.",
		},
	}
	return createDefaultCmd(
		"reddit",
		"Summarizes using Reddit as source",
		"Summarizes using Reddit as source.",
		flags,
	)
}

type aiFlag string

const (
	aiFlagModel  aiFlag = "ai-model"
	aiFlagPrompt aiFlag = "ai-prompt"
)

var defaultAiFlags = []flag{
	{
		Name:        string(aiFlagModel),
		Required:    false,
		Description: "The AI model to use",
	},
	{
		Name:        string(aiFlagPrompt),
		Required:    false,
		Description: "The prompt to use for the AI model",
		Default:     defaultPromptTemplate,
	},
}

const (
	ollamaFlagUrl aiFlag = "url"
)

func ollamaCmd() *cobra.Command {
	flags := []flag{
		{
			Name:        string(ollamaFlagUrl),
			Required:    false,
			Description: "The URl where ollama is accessible",
			Default:     "http://localhost:11434",
		},
	}
	defaultAiFlags[0].Default = "gemma:2b"
	flags = append(flags, defaultAiFlags...)
	cmd := createDefaultCmd(
		"ollama",
		"Summarizes using Ollama AI",
		"Summarizes using Ollama AI.",
		flags,
	)
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		s := sourceByCmd(cmd.Parent())

		aiModel, _ := cmd.Flags().GetString(string(aiFlagModel))
		aiPrompt, _ := cmd.Flags().GetString(string(aiFlagPrompt))
		url, _ := cmd.Flags().GetString(string(ollamaFlagUrl))
		p := &provider.Ollama{
			Common: provider.Common{
				Model:  aiModel,
				Prompt: aiPrompt,
			},
			Url: url,
		}

		return summarize(s, p)
	}
	return cmd
}

const (
	mistralFlagToken aiFlag = "provider-token"
)

func mistralCmd() *cobra.Command {
	flags := []flag{
		{
			Name:        string(mistralFlagToken),
			Required:    true,
			Description: "The API Token for Mistral",
		},
	}
	defaultAiFlags[0].Default = "mistral:7b"
	flags = append(flags, defaultAiFlags...)
	cmd := createDefaultCmd(
		"mistral",
		"Summarizes using Mistral AI",
		"Summarizes using Mistral AI.",
		flags,
	)
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		s := sourceByCmd(cmd.Parent())

		aiModel, _ := cmd.Flags().GetString(string(aiFlagModel))
		aiPrompt, _ := cmd.Flags().GetString(string(aiFlagPrompt))
		apiToken, _ := cmd.Flags().GetString(string(mistralFlagToken))
		p := &provider.Mistral{
			Common: provider.Common{
				Model:  aiModel,
				Prompt: aiPrompt,
			},
			ApiToken: apiToken,
		}

		return summarize(s, p)
	}
	return cmd
}

const (
	openaiFlagToken aiFlag = "provider-token"
)

func openaiCmd() *cobra.Command {
	flags := []flag{
		{
			Name:        string(openaiFlagToken),
			Required:    true,
			Description: "The API Token for OpenAI",
		},
	}
	defaultAiFlags[0].Default = "gpt-3.5-turbo"
	flags = append(flags, defaultAiFlags...)
	cmd := createDefaultCmd(
		"openai",
		"Summarizes using OpenAI AI",
		"Summarizes using OpenAI AI.",
		flags,
	)
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		s := sourceByCmd(cmd.Parent())

		aiModel, _ := cmd.Flags().GetString(string(aiFlagModel))
		aiPrompt, _ := cmd.Flags().GetString(string(aiFlagPrompt))
		apiToken, _ := cmd.Flags().GetString(string(openaiFlagToken))
		p := &provider.OpenAi{
			Common: provider.Common{
				Model:  aiModel,
				Prompt: aiPrompt,
			},
			ApiToken: apiToken,
		}

		return summarize(s, p)
	}
	return cmd
}

func createDefaultCmd(
	name string,
	short string,
	long string,
	flags []flag,
) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   name,
		Short: short,
		Long:  long,
	}

	for _, flag := range flags {
		cmd.PersistentFlags().String(flag.Name, flag.Default, flag.Description)
		if flag.Required {
			cmd.MarkPersistentFlagRequired(flag.Name)
		}
	}

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

func sourceByCmd(cmd *cobra.Command) summaraizer.CommentSource {
	switch cmd.Name() {
	case "github":
		token, _ := cmd.Flags().GetString(string(gitHubFlagSourceToken))

		issue, _ := cmd.Flags().GetString(string(gitHubFlagIssue))
		githubIssueParts := strings.Split(issue, "/")

		return &source.GitHub{
			Token:       token,
			RepoOwner:   githubIssueParts[0],
			RepoName:    githubIssueParts[1],
			IssueNumber: githubIssueParts[2],
		}
	case "reddit":
		post, _ := cmd.Flags().GetString(string(redditFlagPost))

		return &source.Reddit{
			UrlPath: post,
		}
	}
	panic("Unknown source")
}

func summarize(s summaraizer.CommentSource, p summaraizer.AiProvider) error {
	summarization, err := summaraizer.Summarize(s, p)
	if err != nil {
		return err
	}
	fmt.Println(summarization)

	return nil
}
