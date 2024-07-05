package http_monitor

type TaskType uint8

type BaseRequest struct {
	Type string `json:"type"`
}

type TaskRequest struct {
	Type       string `json:"type"`
	UUID       string `json:"uuid"`
	URL        string `json:"url"`
	Timeout    uint   `json:"timeout"`
	TargetAddr string `json:"target_addr,omitempty"`
	CloudApi   Cloud  `json:"cloud_api,omitempty"`
}

type TaskResponse struct {
	Type        string  `json:"type"`
	UUID        string  `json:"uuid"`
	HTTPCode    int     `json:"http_code,omitempty"`
	ErrMsg      string  `json:"err_msg,omitempty"`
	ConsumeTime float64 `json:"consume_time,omitempty"`
	TargetAddr  string  `json:"target_addr,omitempty"`
}

type AuthRequest struct {
	Type  string `json:"type"`
	Token string `json:"token"`
}

type AuthResponse struct {
	Type string `json:"type"`
}

type Cloud struct {
	Provider string `json:"provider"`
	Product  string `json:"product"`
	Region   string `json:"region"`
	Name     string `json:"name"`
}

type CloudKey struct {
	Key    string `json:"key"`
	Secret string `json:"secret"`
}
