package webhook_request

type Request struct {
	Push       Push `json:"push_data"`
	Repository Repo `json:"repository"`
}

type Push struct {
	Digest string `json:"digest"`
	Tag    string `json:"tag"`
}

type Repo struct {
	DateCreated            string `json:"date_created"`
	Name                   string `json:"name"`
	Namespace              string `json:"namespace"`
	Region                 string `json:"region"`
	RepoAuthenticationType string `json:"repo_authentication_type"`
	RepoFullName           string `json:"repo_full_name"`
	RepoOriginType         string `json:"repo_origin_type"`
	RepoType               string `json:"repo_type"`
}
