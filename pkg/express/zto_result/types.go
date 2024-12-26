type zto_result

type AddressInfo struct {
	Name     string `json:"name,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Province string `json:"province,omitempty"`
	City     string `json:"city,omitempty"`
	County   string `json:"county,omitempty"`
	Town     string `json:"town,omitempty"`
	Detail   string `json:"detail,omitempty"`
}

type ReceivingAddress struct {
	ProvinceName  string `json:"provinceName,omitempty"`
	CityName      string `json:"cityName,omitempty"`
	DistrictName  string `json:"districtName,omitempty"`
	DetailAddress string `json:"detailAddress,omitempty"`
}
