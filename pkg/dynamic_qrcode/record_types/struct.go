package record_types

type Request struct {
	Cname string `json:"cname"`
	Path  string `json:"path"`
	IP    string `json:"ip"`
	UA    string `json:"ua"`
}

type Response struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Image string `json:"image"`
}
