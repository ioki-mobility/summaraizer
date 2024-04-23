package provider

import (
	summaraizer "github.com/ioki-mobility/summaraizer/types"
)

type Mistral struct{}

func (m Mistral) Summarize(
	model string,
	comments summaraizer.Comments,
) (string, error) {
	return "This is a summary from mistral", nil
}
