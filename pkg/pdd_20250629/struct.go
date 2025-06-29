package pdd_20250629

type ImageItem struct {
	Url       string `json:"url"`
	HasHidden bool   `json:"has_hidden"`
}

type SpecItem struct {
	Key   string `json:"spec_key"`
	Value string `json:"spec_value"`
}

type SkuItem struct {
	Name   string     `json:"name,omitempty"`
	Image  string     `json:"thumbUrl"`
	Code   string     `json:"out_sku_sn"`
	Suffix string     `json:"suffix,omitempty"`
	Specs  []SpecItem `json:"specs,omitempty"`
}
