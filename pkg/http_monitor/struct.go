package http_monitor

type BaseRequest struct {
	Type string `json:"type"`
}

type TaskItem struct {
	Type 	 string `json:"type"`
	UUID     string `json:"uuid"`
	URL      string `json:"url"`
	HTTPCode int    `json:"http_code"`
	ErrMsg   string `json:"err_msg"`
}

type AuthRequest struct {
	Type string `json:"type"`
	Token string `json:"token"`
}

type AuthResponse struct {
	Type string `json:"type"`
}
