package provider

import (
	"bytes"
	"text/template"

	"github.com/ioki-mobility/summaraizer"
)

// Common is a generic struct that shares common fields between all AI providers.
type Common struct {
	Model  string // The Ai model to use.
	Prompt string // The prompt to use for the AI model.
}

func resolvePrompt(prompt string, comments summaraizer.Comments) (string, error) {
	must := template.Must(template.New("resolvedPrompt").Parse(prompt))
	var buf bytes.Buffer
	err := must.Execute(&buf, comments)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
