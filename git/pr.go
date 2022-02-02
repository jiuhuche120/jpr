package git

import "time"

type PullRequest struct {
	HtmlUrl  string    `json:"html_url"`
	State    string    `json:"state"`
	Title    string    `json:"title"`
	Body     string    `json:"body"`
	CreateAt time.Time `json:"create_at"`
	MergedAt time.Time `json:"merged_at"`
	Head     Head      `json:"head"`
	Base     Base      `json:"base"`
}
type Head struct {
	Ref string `json:"ref"`
}

type Base struct {
	Ref string `json:"ref"`
}
