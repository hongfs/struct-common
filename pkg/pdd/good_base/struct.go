package good_base

type AutoGenerated struct {
	InitDataObj InitDataObj `json:"initDataObj,omitempty"`
}

type TopGallery struct {
	URL string `json:"url,omitempty"`
}

type DetailGallery struct {
	URL    string `json:"url,omitempty"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}

type Specs struct {
	SpecKey     string `json:"spec_key,omitempty"`
	SpecValue   string `json:"spec_value,omitempty"`
	SpecKeyID   int    `json:"spec_key_id,omitempty"`
	SpecValueID int64  `json:"spec_value_id,omitempty"`
}

type Skus struct {
	SkuID        int64   `json:"skuId,omitempty"`
	GoodsID      int64   `json:"goodsId,omitempty"`
	ThumbURL     string  `json:"thumbUrl,omitempty"`
	InitQuantity int     `json:"initQuantity,omitempty"`
	Quantity     int     `json:"quantity,omitempty"`
	Spec         string  `json:"spec,omitempty"`
	Specs        []Specs `json:"specs,omitempty"`
	Price        int     `json:"price,omitempty"` // skuPrice
	SkuPrice     int     `json:"skuPrice,omitempty"`
	NormalPrice  float64 `json:"normalPrice,omitempty"`
	GroupPrice   float64 `json:"groupPrice,omitempty"`
}

type GoodsProperty struct {
	Key    string   `json:"key,omitempty"`
	Values []string `json:"values,omitempty"`
	RefPid int      `json:"ref_pid,omitempty"`
}

type Goods struct {
	CatID         int             `json:"catID,omitempty"`
	GoodsID       int64           `json:"goodsID,omitempty"`
	GoodsName     string          `json:"goodsName,omitempty"`
	ShareDesc     string          `json:"shareDesc,omitempty"`
	GoodsType     int             `json:"goodsType,omitempty"`
	TopGallery    []TopGallery    `json:"topGallery,omitempty"`
	DetailGallery []DetailGallery `json:"detailGallery,omitempty"`
	Skus          []Skus          `json:"skus,omitempty"`
	ThumbURL      string          `json:"thumbUrl,omitempty"`
	HdThumbURL    string          `json:"hdThumbUrl,omitempty"`
	CatID1        int             `json:"catID1,omitempty"`
	CatID2        int             `json:"catID2,omitempty"`
	CatID3        int             `json:"catID3,omitempty"`
	CatID4        int             `json:"catID4"`
	GoodsProperty []GoodsProperty `json:"goodsProperty,omitempty"`
	LinePrice     string          `json:"linePrice,omitempty"`
}

type InitDataObj struct {
	Goods Goods `json:"goods,omitempty"`
}
