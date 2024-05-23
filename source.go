package summaraizer

import (
	"encoding/json"
	"io"
)

// CommentSource is an interface that defines a source to fetch comments from.
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
