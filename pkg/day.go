package pkg

type Day struct {
	Code    int     `json:"code,omitempty"`
	Type    Type    `json:"type"`
	Holiday Holiday `json:"holiday"`
}

type Type struct {
	Type int    `json:"type,omitempty"`
	Name string `json:"name,omitempty"`
	Week int    `json:"week,omitempty"`
}

type Holiday struct {
	Holiday bool   `json:"holiday,omitempty"`
	Name    string `json:"name,omitempty"`
	Wage    int    `json:"wage,omitempty"`
	After   bool   `json:"after,omitempty"`
	Target  string `json:"target,omitempty"`
}
