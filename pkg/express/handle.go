package express

type Trace struct {
	Area   string `json:"area"`
	Desc   string `json:"desc"`
	Time   int64  `json:"time"`
	Status string `json:"status"`
}

type Data struct {
	CpCode      string  `json:"cpCode"`
	CpMobile    string  `json:"cpMobile"`
	CpName      string  `json:"cpName"`
	No          string  `json:"no"`
	LastMessage string  `json:"lastMessage"`
	Status      string  `json:"status"`
	Traces      []Trace `json:"traces"`
}
