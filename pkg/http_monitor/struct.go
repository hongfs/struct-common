package http_monitor

type BaseRequest struct {
	Type string `json:"type"`
}

type TaskRequest struct {
	Type     string `json:"type"`
	UUID     string `json:"uuid"`
	URL      string `json:"url"`
	Timeout  uint   `json:"timeout"`
	TargetIP string `json:"target_ip,omitempty"`
}

type TaskResponse struct {
	Type        string  `json:"type"`
	UUID        string  `json:"uuid"`
	HTTPCode    int     `json:"http_code,omitempty"`
	ErrMsg      string  `json:"err_msg,omitempty"`
	ConsumeTime float64 `json:"consume_time,omitempty"`
}

type AuthRequest struct {
	Type  string `json:"type"`
	Token string `json:"token"`
}

type AuthResponse struct {
	Type string `json:"type"`
}
