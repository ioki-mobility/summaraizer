package provider

import (
	summaraizer "github.com/ioki-mobility/summaraizer/types"
)

type OpenAi struct{}

func (o OpenAi) Summarize(
	model string,
	comments summaraizer.Comments,
) (string, error) {
	return "This is a summary from OpenAi", nil
}
