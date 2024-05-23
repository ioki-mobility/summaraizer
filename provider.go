package summaraizer

import (
	"bytes"
	"encoding/json"
	"io"
	"text/template"
)

// AiProvider is an interface that defines an AI provider to summarize comments.
type AiProvider interface {
	Summarize(reader io.Reader) (string, error)
}

// Common is a generic struct that shares common fields between all AI providers.
type Common struct {
	Model  string // The Ai model to use.
	Prompt string // The prompt to use for the AI model.
}

func decodeAndSummarize(reader io.Reader, f func(comments Comments) (string, error)) (string, error) {
	var comments Comments
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&comments); err != nil {
		return "", err
	}
	return f(comments)
}

func resolvePrompt(prompt string, comments Comments) (string, error) {
	must := template.Must(template.New("resolvedPrompt").Parse(prompt))
	var buf bytes.Buffer
	err := must.Execute(&buf, comments)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
