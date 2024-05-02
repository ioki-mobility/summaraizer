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
		url, _ := cmd.Flags().GetString("url")
		provider := &provider.Ollama{
			Model: aiModel,
			Url:   url,
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
		apiToken, _ := cmd.Flags().GetString("api-token")
		provider := &provider.Mistral{
			Model:    aiModel,
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
		apiToken, _ := cmd.Flags().GetString("api-token")
		provider := &provider.OpenAi{
			Model:    aiModel,
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

	ollamaCmd.PersistentFlags().String("ai-model", "mistral:7b", "AI Model")
	ollamaCmd.PersistentFlags().String("url", "http://localhost:11434", "The URl where ollama is accessible")
	rootCmd.AddCommand(ollamaCmd)

	mistralCmd.PersistentFlags().String("ai-model", "mistral-small-latest", "AI Model")
	mistralCmd.PersistentFlags().String("api-token", "", "The API Token for Mistral")
	mistralCmd.MarkPersistentFlagRequired("api-token")
	rootCmd.AddCommand(mistralCmd)

	openaiCmd.PersistentFlags().String("ai-model", "gpt-3.5-turbo", "AI Model")
	openaiCmd.PersistentFlags().String("api-token", "", "The API Token for OpenAI")
	openaiCmd.MarkPersistentFlagRequired("api-token")
	rootCmd.AddCommand(openaiCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
