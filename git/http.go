package git

import "net/http"

type ClientOption = func(http.RoundTripper) http.RoundTripper

func NewHttpClient(opts ...ClientOption) *http.Client {
	tr := http.DefaultTransport
	for _, opt := range opts {
		tr = opt(tr)
	}
	return &http.Client{Transport: tr}
}

type funcTripper struct {
	roundTrip func(*http.Request) (*http.Response, error)
}

func (tr funcTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return tr.roundTrip(req)
}

func AddHeader(name, value string) ClientOption {
	return func(tr http.RoundTripper) http.RoundTripper {
		return &funcTripper{roundTrip: func(req *http.Request) (*http.Response, error) {
			if req.Header.Get(name) == "" {
				req.Header.Add(name, value)
			}
			return tr.RoundTrip(req)
		}}
	}
}


