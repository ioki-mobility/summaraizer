package cli

import (
	"os"
	"strings"

	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type config struct {
	GitHub gitHubConfig
	GitLab gitLabConfig
	Slack  slackConfig

	Ollama    ollamaConfig
	OpenAI    openAiConfig
	Anthropic anthropicConfig
	Google    googleConfig
}

type gitHubConfig struct {
	Token string
}

type gitLabConfig struct {
	Token string
	Url   string
}

type slackConfig struct {
	Token string
}

type ollamaConfig struct {
	Url    string
	Model  string
	Prompt string
}

type openAiConfig struct {
	Token  string
	Model  string
	Prompt string
}

type anthropicConfig struct {
	Token  string
	Model  string
	Prompt string
}

type googleConfig struct {
	Model  string
	Token  string
	Prompt string
}

func newConfig() *config {
	k := koanf.New(".")
	k.Load(file.Provider("config.json"), json.Parser())
	k.Load(file.Provider(os.ExpandEnv("$HOME/.config/summaraizer/config.json")), json.Parser())
	k.Load(file.Provider(os.ExpandEnv("$HOME/.summaraizer/config.json")), json.Parser())
	k.Load(env.Provider("SUMMARAIZER_", ".", func(s string) string {
		envWithoutPrefix := strings.TrimPrefix(s, "SUMMARAIZER_")
		return strings.Replace(strings.ToLower(envWithoutPrefix), "_", ".", -1)
	}), nil)
	var cfg config
	if err := k.Unmarshal("", &cfg); err != nil {
		return nil
	}
	return &cfg
}
