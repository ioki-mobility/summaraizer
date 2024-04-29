package main

import (
	"fmt"
	"os"

	summaraizer "github.com/ioki-mobility/summaraizer"
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
		aiModel, _ := cmd.Flags().GetString("ai-model")
		token, _ := cmd.Flags().GetString("token")
		owner, _ := cmd.Flags().GetString("owner")
		repo, _ := cmd.Flags().GetString("repo")
		issueNumber, _ := cmd.Flags().GetString("issue-number")

		input := summaraizer.Input{
			AiInput: summaraizer.AiInput{
				AiProviderName: summaraizer.Ollama,
				AiModel:        aiModel,
			},
			GitHubInput: summaraizer.GitHubInput{
				Token:       token,
				RepoOwner:   owner,
				RepoName:    repo,
				IssueNumber: issueNumber,
			},
		}
		summarization, err := summaraizer.Summarize(input)
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
		aiModel, _ := cmd.Flags().GetString("ai-model")
		token, _ := cmd.Flags().GetString("token")
		owner, _ := cmd.Flags().GetString("owner")
		repo, _ := cmd.Flags().GetString("repo")
		issueNumber, _ := cmd.Flags().GetString("issue-number")

		input := summaraizer.Input{
			AiInput: summaraizer.AiInput{
				AiProviderName: summaraizer.Mistral,
				AiModel:        aiModel,
			},
			GitHubInput: summaraizer.GitHubInput{
				Token:       token,
				RepoOwner:   owner,
				RepoName:    repo,
				IssueNumber: issueNumber,
			},
		}
		summarization, err := summaraizer.Summarize(input)
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
		aiModel, _ := cmd.Flags().GetString("ai-model")
		token, _ := cmd.Flags().GetString("token")
		owner, _ := cmd.Flags().GetString("owner")
		repo, _ := cmd.Flags().GetString("repo")
		issueNumber, _ := cmd.Flags().GetString("issue-number")

		input := summaraizer.Input{
			AiInput: summaraizer.AiInput{
				AiProviderName: summaraizer.OpenAI,
				AiModel:        aiModel,
			},
			GitHubInput: summaraizer.GitHubInput{
				Token:       token,
				RepoOwner:   owner,
				RepoName:    repo,
				IssueNumber: issueNumber,
			},
		}
		summarization, err := summaraizer.Summarize(input)
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

	ollamaCmd.PersistentFlags().String("ai-model", "mistral:7b", "AI Model")
	rootCmd.AddCommand(ollamaCmd)

	mistralCmd.PersistentFlags().String("ai-model", "mistral-small-latest", "AI Model")
	rootCmd.AddCommand(mistralCmd)

	openaiCmd.PersistentFlags().String("ai-model", "gpt-3.5-turbo", "AI Model")
	rootCmd.AddCommand(openaiCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
