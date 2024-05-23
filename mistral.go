package summaraizer

import (
	"fmt"
	"io"
)

// Mistral is a provider that uses the Mistral API as an AI provider.
type Mistral struct {
	Common
	ApiToken string
}

func (m *Mistral) Summarize(reader io.Reader) (string, error) {
	return decodeAndSummarize(reader, func(comments Comments) (string, error) {
		return fmt.Sprintf("This is a summary from mistral with model %s", m.Model), nil
	})
}
