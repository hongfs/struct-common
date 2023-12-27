package github_status

type Incident struct {
	Status string `json:"status"`
}

type Component struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type Response struct {
	Incidents  []Incident  `json:"incidents"`
	Components []Component `json:"components"`

	Status struct {
		Incident string `json:"indicator"`
	}
}
