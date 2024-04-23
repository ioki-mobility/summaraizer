package github

import "net/http"

func newRequest(method, url string, token string) http.Request {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}
	addHeaders(req, token)
	return *req
}

func addHeaders(req *http.Request, token string) {
	req.Header.Add("Accept", "application/vnd.github+json")
	// Public repos doesn't need Authorization header
	if token != "" {
		req.Header.Add("Authorization", "Bearer "+token)
	}
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")
}
