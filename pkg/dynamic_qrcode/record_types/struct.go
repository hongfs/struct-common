package record_types

type Request struct {
	Cname string `json:"cname"`
	Path  string `json:"path"`
	IP    string `json:"ip"`
	UA    string `json:"ua"`
}

type Response struct {
	Title string `json:"title"`
	Image string `json:"image"`
}
