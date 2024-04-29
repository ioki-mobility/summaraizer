package provider

import (
	"fmt"

	summaraizer "github.com/ioki-mobility/summaraizer/types"
)

type OpenAi struct {
	Model    string
	ApiToken string
}

func (o OpenAi) Summarize(
	comments summaraizer.Comments,
) (string, error) {
	return fmt.Sprintf("This is a summary from OpenAI with model %s", o.Model), nil
}
