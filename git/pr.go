package git

import "time"

type PullRequest struct {
	ID       int64     `json:"id"`
	Url      string    `json:"url"`
	HtmlUrl  string    `json:"html_url"`
	State    string    `json:"state"`
	Title    string    `json:"title"`
	User     User      `json:"user"`
	Body     string    `json:"body"`
	CreateAt time.Time `json:"create_at"`
	MergedAt time.Time `json:"merged_at"`
	Head     Head      `json:"head"`
	Base     Base      `json:"base"`
	DingTalk string
}
type Head struct {
	Ref string `json:"ref"`
}

type Base struct {
	Ref string `json:"ref"`
}

type User struct {
	Login string `json:"login"`
}
