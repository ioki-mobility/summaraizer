package provider

import (
	"fmt"

	"github.com/ioki-mobility/summaraizer"
)

// Mistral is a provider that uses the Mistral API as an AI provider.
type Mistral struct {
	Common
	ApiToken string
}

func (m *Mistral) Summarize(
	comments summaraizer.Comments,
) (string, error) {
	return fmt.Sprintf("This is a summary from mistral with model %s", m.Model), nil
}
