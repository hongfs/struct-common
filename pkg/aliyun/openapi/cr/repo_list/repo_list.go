package repo_list

type RepoDomainList struct {
	Internal string `json:"internal,omitempty"`
	Public   string `json:"public,omitempty"`
	Vpc      string `json:"vpc,omitempty"`
}

type Repo struct {
	Summary           string         `json:"summary,omitempty"`
	RepoID            int            `json:"repoId,omitempty"`
	GmtModified       int64          `json:"gmtModified,omitempty"`
	RepoNamespace     string         `json:"repoNamespace,omitempty"`
	RepoName          string         `json:"repoName,omitempty"`
	RepoOriginType    string         `json:"repoOriginType,omitempty"`
	Stars             int            `json:"stars,omitempty"`
	GmtCreate         int64          `json:"gmtCreate,omitempty"`
	RepoBuildType     string         `json:"repoBuildType,omitempty"`
	RepoType          string         `json:"repoType,omitempty"`
	RepoDomainList    RepoDomainList `json:"repoDomainList,omitempty"`
	Downloads         int            `json:"downloads,omitempty"`
	RegionID          string         `json:"regionId,omitempty"`
	Logo              string         `json:"logo,omitempty"`
	RepoStatus        string         `json:"repoStatus,omitempty"`
	RepoAuthorizeType string         `json:"repoAuthorizeType,omitempty"`
}

type Data struct {
	Total    int    `json:"total,omitempty"`
	Repos    []Repo `json:"repos,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
	Page     int    `json:"page,omitempty"`
}

type ResponseData struct {
	Data Data `json:"data,omitempty"`
}
