package douyin_parse

type Response struct {
	Success  bool     `json:"success"`
	Video    string   `json:"video"`
	Duration int      `json:"duration"`
	Cover    string   `json:"cover"`
	Images   []string `json:"images"`
	Msg      string   `json:"msg"`
	Desc     string   `json:"desc"`
}
