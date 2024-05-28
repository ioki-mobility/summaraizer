package summaraizer

import (
	"encoding/json"
	"io"
)

// CommentSource fetches comments from a source and writes them to a writer.
type CommentSource interface {
	Fetch(writer io.Writer) error
}

func fetchAndEncode(writer io.Writer, fetch func() (Comments, error)) error {
	comments, err := fetch()
	if err != nil {
		return err

	}
	encoder := json.NewEncoder(writer)
	return encoder.Encode(comments)
}
