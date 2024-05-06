package main

import (
	"fmt"
	"os"

	"github.com/ioki-mobility/summaraizer"
	"github.com/ioki-mobility/summaraizer/provider"
	"github.com/ioki-mobility/summaraizer/source"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "summaraizer",
	Short: "Summarizes GitHub issue comments",
	Long:  `A tool to summarize GitHub issue (or pull request) comments using AI.`,
}

var ollamaCmd = &cobra.Command{
	Use:   "ollama",
	Short: "Summarizes using Ollama AI",
	Long:  `Summarizes using Ollama AI.`,
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		owner, _ := cmd.Flags().GetString("owner")
		repo, _ := cmd.Flags().GetString("repo")
		issueNumber, _ := cmd.Flags().GetString("issue-number")
		gitHubInput := &source.GitHub{
			Token:       token,
			RepoOwner:   owner,
			RepoName:    repo,
			IssueNumber: issueNumber,
		}

		aiModel, _ := cmd.Flags().GetString("ai-model")
		aiPrompt, _ := cmd.Flags().GetString("ai-prompt")
		url, _ := cmd.Flags().GetString("url")
		provider := &provider.Ollama{
			Common: provider.Common{
				Model:  aiModel,
				Prompt: aiPrompt,
			},
			Url: url,
		}

		summarization, err := summaraizer.Summarize(gitHubInput, provider)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println(summarization)
	},
}

var mistralCmd = &cobra.Command{
	Use:   "mistral",
	Short: "Summarizes using Mistral AI",
	Long:  `Summarizes using Mistral AI.`,
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		owner, _ := cmd.Flags().GetString("owner")
		repo, _ := cmd.Flags().GetString("repo")
		issueNumber, _ := cmd.Flags().GetString("issue-number")
		gitHub := &source.GitHub{
			Token:       token,
			RepoOwner:   owner,
			RepoName:    repo,
			IssueNumber: issueNumber,
		}

		aiModel, _ := cmd.Flags().GetString("ai-model")
		aiPrompt, _ := cmd.Flags().GetString("ai-prompt")
		apiToken, _ := cmd.Flags().GetString("api-token")
		provider := &provider.Mistral{
			Common: provider.Common{
				Model:  aiModel,
				Prompt: aiPrompt,
			},
			ApiToken: apiToken,
		}

		summarization, err := summaraizer.Summarize(gitHub, provider)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println(summarization)
	},
}

var openaiCmd = &cobra.Command{
	Use:   "openai",
	Short: "Summarizes using OpenAI AI",
	Long:  `Summarizes using OpenAI AI.`,
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		owner, _ := cmd.Flags().GetString("owner")
		repo, _ := cmd.Flags().GetString("repo")
		issueNumber, _ := cmd.Flags().GetString("issue-number")
		gitHubInput := &source.GitHub{
			Token:       token,
			RepoOwner:   owner,
			RepoName:    repo,
			IssueNumber: issueNumber,
		}

		aiModel, _ := cmd.Flags().GetString("ai-model")
		aiPrompt, _ := cmd.Flags().GetString("ai-prompt")
		apiToken, _ := cmd.Flags().GetString("api-token")
		provider := &provider.OpenAi{
			Common: provider.Common{
				Model:  aiModel,
				Prompt: aiPrompt,
			},
			ApiToken: apiToken,
		}

		summarization, err := summaraizer.Summarize(gitHubInput, provider)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println(summarization)
	},
}

func main() {
	rootCmd.PersistentFlags().String("token", "", "GitHub Token (can be empty for public repositories)")
	rootCmd.PersistentFlags().String("owner", "", "GitHub Repository Owner")
	rootCmd.PersistentFlags().String("repo", "", "GitHub Repository Name")
	rootCmd.PersistentFlags().String("issue-number", "", "GitHub Issue Number")
	rootCmd.MarkPersistentFlagRequired("owner")
	rootCmd.MarkPersistentFlagRequired("repo")
	rootCmd.MarkPersistentFlagRequired("issue-number")

	createDefaultAiFlags(ollamaCmd, "mistral:7b", defaultPromptTemplate)
	ollamaCmd.PersistentFlags().String("url", "http://localhost:11434", "The URl where ollama is accessible")
	rootCmd.AddCommand(ollamaCmd)

	createDefaultAiFlags(mistralCmd, "mistral:7b", defaultPromptTemplate)
	mistralCmd.PersistentFlags().String("api-token", "", "The API Token for Mistral")
	mistralCmd.MarkPersistentFlagRequired("api-token")
	rootCmd.AddCommand(mistralCmd)

	createDefaultAiFlags(openaiCmd, "gpt-3.5-turbo", defaultPromptTemplate)
	openaiCmd.PersistentFlags().String("api-token", "", "The API Token for OpenAI")
	openaiCmd.MarkPersistentFlagRequired("api-token")
	rootCmd.AddCommand(openaiCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func createDefaultAiFlags(cmd *cobra.Command, model string, prompt string) {
	cmd.PersistentFlags().String("ai-model", model, "AI Model")
	cmd.PersistentFlags().String("ai-prompt", prompt, "The prompt to use for the AI model")
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
