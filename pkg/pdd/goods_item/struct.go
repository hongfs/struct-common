package goods_item

type Spec struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

type Property struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
	Pid    int64    `json:"ref_pid"`
}
