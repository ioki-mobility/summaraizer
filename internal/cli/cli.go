package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/ioki-mobility/summaraizer"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	conf := newConfig()
	var rootCmd = &cobra.Command{
		Use:   "summaraizer",
		Short: "Summarizes comments",
		Long:  "Summarizes comments from a variety of sources using AI models from different providers.",
	}
	rootCmd.AddCommand(githubCmd(conf), redditCmd(conf), gitlabCmd(conf), slackCmd(conf))
	rootCmd.AddCommand(ollamaCmd(conf), openaiCmd(conf), anthropicCmd(conf), googleCmd(conf))
	return rootCmd
}

func githubCmd(c *config) *cobra.Command {
	flagIssue := "issue"
	flagToken := "token"
	var cmd = &cobra.Command{
		Use:   "github",
		Short: "Fetches comments from GitHub",
		Long: `To summarize a GitHub issue, use the format owner/repo/issue_number.
At GitHub terminology, a pull request is also an issue. Therefore, you can summarize a pull request using the same format.
If the repository is private, you need to provide a token.`,
		Example: `summaraizer github --issue ioki-mobility/summaraizer/1
summaraizer github --issue ioki-mobility/summaraizer/1 --token <token>`,
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
	if tokenConfig := c.GetString("github.token"); tokenConfig != "" {
		cmd.Flags().Set(flagToken, tokenConfig)
	}

	return cmd
}

func redditCmd(c *config) *cobra.Command {
	flagPost := "post"
	cmd := &cobra.Command{
		Use:   "reddit",
		Short: "Fetches comments from a Reddit post",
		Long: `To summarize a Reddit post, use the URL path. Everything after reddit.com without the leading slash.
Note that we only fetch the top-level comments. Nested comments are ignored.`,
		Example: "summaraizer reddit --post r/ArtificialInteligence/comments/1d16cxl/miss_ai_worlds_first_beauty_contest_with_computer/",
		RunE: func(cmd *cobra.Command, args []string) error {
			post, _ := cmd.Flags().GetString(flagPost)

			s := &summaraizer.Reddit{
				UrlPath: post,
			}
			return fetch(s)
		},
	}

	cmd.Flags().String(flagPost, "", "The Reddit post to summarize. Use the URL path.")
	cmd.MarkFlagRequired(flagPost)

	return cmd
}

func gitlabCmd(c *config) *cobra.Command {
	flagIssue := "issue"
	flagToken := "token"
	flagUrl := "url"
	var cmd = &cobra.Command{
		Use:   "gitlab",
		Short: "Fetches comments from GitLab issues",
		Long: `To summarize a GitLab issue, use the format owner/repo/issue_number. 
You always have to provide a token.
If you have a custom GitLab instance, you can provide the URL with the --url flag.
Note that we only fetch the top-level comments. Nested comments are ignored.`,
		Example: `summaraizer gitlab --issue ioki-mobility/summaraizer/1 --token <token>
summaraizer gitlab --issue ioki-mobility/summaraizer/1 --token <token> --url https://gitlab.url.com`,
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
	if tokenConfig := c.GetString("gitlab.token"); tokenConfig != "" {
		cmd.Flags().Set(flagToken, tokenConfig)
	}

	cmd.Flags().String(flagUrl, "https://gitlab.com", "The URL of the GitLab instance.")
	if urlConfig := c.GetString("gitlab.url"); urlConfig != "" {
		cmd.Flags().Set(flagUrl, urlConfig)
	}

	return cmd
}

func slackCmd(c *config) *cobra.Command {
	flagToken := "token"
	flagChannel := "channel"
	flagTs := "ts"
	var cmd = &cobra.Command{
		Use:   "slack",
		Short: "Fetches comments from a Slack thread",
		Long: `To summarize a Slack thread, you need to provide the token, the channel ID, and the timestamp of the thread.
You can get the channel ID and the timestamp from the URL of the thread.`,
		Example: "summaraizer slack --token <token> --channel <channel_id> --ts <timestamp>",
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
	if tokenConfig := c.GetString("slack.token"); tokenConfig != "" {
		cmd.Flags().Set(flagToken, tokenConfig)
	}

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

func ollamaCmd(c *config) *cobra.Command {
	flagUrl := "url"
	var cmd = &cobra.Command{
		Use:   "ollama",
		Short: "Summarizes using Ollama",
		Long: `To summarize using Ollama, you *can* provide the URL where Ollama is accessible.
If you are running Ollama locally, you can use the default URL. There is no need to provide the URL.
Optional flags are the AI model and the prompt. The prompt can make use of Go template functions.`,
		Example: `summaraizer ollama
summaraizer ollama --model llama3.2:3b`,
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
	if modelConfig := c.GetString("ollama.model"); modelConfig != "" {
		cmd.Flags().Set(aiFlagModel, modelConfig)
	}

	cmd.Flags().String(aiFlagPrompt, defaultPromptTemplate, "The prompt to use for the AI model")
	if promptConfig := c.GetString("ollama.prompt"); promptConfig != "" {
		cmd.Flags().Set(aiFlagPrompt, promptConfig)
	}

	cmd.Flags().String(flagUrl, "http://localhost:11434", "The URl where ollama is accessible")
	if urlConfig := c.GetString("ollama.url"); urlConfig != "" {
		cmd.Flags().Set(flagUrl, urlConfig)
	}

	return cmd
}

func openaiCmd(c *config) *cobra.Command {
	flagToken := "token"
	cmd := &cobra.Command{
		Use:   "openai",
		Short: "Summarizes using OpenAI",
		Long: `So summarize using OpenAI, you need to provide the API token.
Optional flags are the AI model and the prompt. The prompt can make use of Go template functions.`,
		Example: `summaraizer openai --token <token>
summaraizer openai --token <token> --model gpt4o-mini`,
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
	if modelConfig := c.GetString("openai.model"); modelConfig != "" {
		cmd.Flags().Set(aiFlagModel, modelConfig)
	}

	cmd.Flags().String(aiFlagPrompt, defaultPromptTemplate, "The prompt to use for the AI model")
	if promptConfig := c.GetString("openai.prompt"); promptConfig != "" {
		cmd.Flags().Set(aiFlagPrompt, promptConfig)
	}

	cmd.Flags().String(flagToken, "", "The API Token for OpenAI")
	cmd.MarkFlagRequired(flagToken)
	if tokenConfig := c.GetString("openai.token"); tokenConfig != "" {
		cmd.Flags().Set(flagToken, tokenConfig)
	}

	return cmd
}

func anthropicCmd(c *config) *cobra.Command {
	flagToken := "token"
	cmd := &cobra.Command{
		Use:   "anthropic",
		Short: "Summarizes using Anthropic",
		Long: `To summarize using Anthropic, you need to provide the API token.
Optional flags are the AI model and the prompt. The prompt can make use of Go template functions.`,
		Example: `summaraizer anthropic --token <token>
summaraizer anthropic --token <token> --model claude-3-5-sonnet-20241022`,
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
	if modelConfig := c.GetString("anthropic.model"); modelConfig != "" {
		cmd.Flags().Set(aiFlagModel, modelConfig)
	}

	cmd.Flags().String(aiFlagPrompt, defaultPromptTemplate, "The prompt to use for the AI model")
	if promptConfig := c.GetString("anthropic.prompt"); promptConfig != "" {
		cmd.Flags().Set(aiFlagPrompt, promptConfig)
	}

	cmd.Flags().String(flagToken, "", "The API Token for Anthropic")
	cmd.MarkFlagRequired(flagToken)
	if tokenConfig := c.GetString("anthropic.token"); tokenConfig != "" {
		cmd.Flags().Set(flagToken, tokenConfig)
	}

	return cmd
}

func googleCmd(c *config) *cobra.Command {
	flagToken := "token"
	cmd := &cobra.Command{
		Use:   "google",
		Short: "Summarizes using Google",
		Long: `To summarize using Google, you need to provide the API token.
Optional flags are the AI model and the prompt. The prompt can make use of Go template functions.`,
		Example: `summaraizer google --token <token>
summaraizer google --token <token> --model gemini-1.5-pro-002`,
		RunE: func(cmd *cobra.Command, args []string) error {
			aiModel, _ := cmd.Flags().GetString(aiFlagModel)
			aiPrompt, _ := cmd.Flags().GetString(aiFlagPrompt)
			apiToken, _ := cmd.Flags().GetString(flagToken)

			p := &summaraizer.Google{
				Model:    aiModel,
				Prompt:   aiPrompt,
				ApiToken: apiToken,
			}

			return summarize(p)
		},
	}

	cmd.Flags().String(aiFlagModel, "gemini-1.5-flash-8b", "The AI model to use")
	if modelConfig := c.GetString("google.model"); modelConfig != "" {
		cmd.Flags().Set(aiFlagModel, modelConfig)
	}

	cmd.Flags().String(aiFlagPrompt, defaultPromptTemplate, "The prompt to use for the AI model")
	if promptConfig := c.GetString("google.prompt"); promptConfig != "" {
		cmd.Flags().Set(aiFlagPrompt, promptConfig)
	}

	cmd.Flags().String(flagToken, "", "The API Token for Google")
	cmd.MarkFlagRequired(flagToken)
	if tokenConfig := c.GetString("google.token"); tokenConfig != "" {
		cmd.Flags().Set(flagToken, tokenConfig)
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
