package short_url

type ValueType int

const (
	PermanentUrl  ValueType = 1 // 永久 URL
	GeneralUrl    ValueType = 2 // 通用 URL
	ShellUrl      ValueType = 3
	MarkUrl       ValueType = 4
	QRCodeContent ValueType = 11 // 内容营销
	InvalidID     ValueType = 99
)

type Value struct {
	Type        ValueType `json:"type,omitempty"`
	ID          int64     `json:"id,omitempty"`
	Value       string    `json:"value,omitempty"`
	UA          string    `json:"ua,omitempty"`
	Mark        string    `json:"mark,omitempty"`
	EnableQuota bool      `json:"enable_quota,omitempty"`
	AccessQuota int64     `json:"access_quota,omitempty"`
}
