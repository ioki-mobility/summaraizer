package provider

import (
	"fmt"

	"github.com/ioki-mobility/summaraizer"
)

type OpenAi struct {
	Model    string
	ApiToken string
}

func (o *OpenAi) Summarize(
	comments summaraizer.Comments,
) (string, error) {
	return fmt.Sprintf("This is a summary from OpenAI with model %s", o.Model), nil
}
