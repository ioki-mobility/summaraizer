package provider

import (
	"fmt"

	"github.com/ioki-mobility/summaraizer"
)

// OpenAi is a provider that uses OpenAI API as an AI provider.
type OpenAi struct {
	Common
	ApiToken string
}

func (o *OpenAi) Summarize(
	comments summaraizer.Comments,
) (string, error) {
	return fmt.Sprintf("This is a summary from OpenAI with model %s", o.Model), nil
}
