package provider

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/ioki-mobility/summaraizer"
)

// Mistral is a provider that uses the Mistral API as an AI provider.
type Mistral struct {
	Common
	ApiToken string
}

func (m *Mistral) Summarize(reader io.Reader) (string, error) {
	var comments summaraizer.Comments
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&comments); err != nil {
		return "", err
	}
	return m.summarizeInternal(comments)
}

func (m *Mistral) summarizeInternal(
	comments summaraizer.Comments,
) (string, error) {
	return fmt.Sprintf("This is a summary from mistral with model %s", m.Model), nil
}
