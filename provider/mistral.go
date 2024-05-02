package provider

import (
	"fmt"

	"github.com/ioki-mobility/summaraizer"
)

type Mistral struct {
	Model    string
	ApiToken string
}

func (m *Mistral) Summarize(
	comments summaraizer.Comments,
) (string, error) {
	return fmt.Sprintf("This is a summary from mistral with model %s", m.Model), nil
}
