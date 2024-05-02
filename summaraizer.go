package summaraizer

type Comments []Comment

type Comment struct {
	Author string
	Body   string
}

type CommentSource interface {
	Fetch() (Comments, error)
}

type AiProvider interface {
	Summarize(comments Comments) (string, error)
}

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
