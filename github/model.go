package github

type Repo struct {
	Owner string
	Name  string
}

type Issue struct {
	Repo
	Number string
}

type IssueResponse struct {
	Body     string       `json:"body"`
	User     UserResponse `json:"user"`
	Comments int          `json:"comments"`
}

type UserResponse struct {
	Login string `json:"login"`
}

type IssueCommentsResponse = []IssueCommentResponse

type IssueCommentResponse struct {
	Body string       `json:"body"`
	User UserResponse `json:"user"`
}
