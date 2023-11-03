package http_monitor

type TaskItem struct {
	UUID     string `json:"uuid"`
	URL      string `json:"url"`
	HTTPCode int    `json:"http_code"`
	ErrMsg   string `json:"err_msg"`
}
