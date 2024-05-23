package summaraizer

import "io"

type Comments []Comment

type Comment struct {
	Author string
	Body   string
}

// CommentSource is an interface that defines a source to fetch comments from.
type CommentSource interface {
	Fetch(writer io.Writer) error
}

// AiProvider is an interface that defines an AI provider to summarize comments.
type AiProvider interface {
	Summarize(reader io.Reader) (string, error)
}
