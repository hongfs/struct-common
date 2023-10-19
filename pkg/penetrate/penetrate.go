package penetrate

import "net/http"

type BaseRequest struct {
	Type string `json:"type"`
}

type RegisterResponse struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

type BaseApiResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type ReceiveResponse struct {
	Type       string      `json:"type"`
	UUID       string      `json:"uuid"`
	HttpStatus int         `json:"http_status"`
	Header     http.Header `json:"header"`
	Body       []byte      `json:"body"`
}

type ForwardRequest struct {
	Type   string      `json:"type"`
	UUID   string      `json:"uuid"`
	Method string      `json:"method"`
	Url    string      `json:"url"`
	Header http.Header `json:"header"`
	Body   []byte      `json:"body"`
}

type AuthRequest struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Token string `json:"token"`
}
