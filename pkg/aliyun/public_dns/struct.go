package public_dns

type Answer struct {
	Name string `json:"name"`
	Type int    `json:"type"`
	Data string `json:"data"`
}

type Response struct {
	Status int      `json:"Status"`
	Answer []Answer `json:"Answer"`
}
