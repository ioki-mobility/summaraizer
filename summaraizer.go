package summaraizer

type Comments []Comment

type Comment struct {
	Author string
	Body   string
}

// CommentSource is an interface that defines a source to fetch comments from.
type CommentSource interface {
	Fetch() (Comments, error)
}

// AiProvider is an interface that defines an AI provider to summarize comments.
type AiProvider interface {
	Summarize(comments Comments) (string, error)
}

// Summarize fetches comments from a source and summarizes them using an AI provider.
func Summarize(source CommentSource, aiProvider AiProvider) (string, error) {
	comments, err := source.Fetch()
	if err != nil {
		return "", err
	}

	summarization, err := aiProvider.Summarize(comments)
	if err != nil {
		return "", err
	}

	return summarization, nil
}
