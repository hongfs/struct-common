package package_event

import "github.com/google/go-github/v56/github"

type PackageVersion struct {
	InstallationCommand *string `json:"installation_command,omitempty"`
}

type Package struct {
	PackageVersion *PackageVersion `json:"package_version,omitempty"`
}

type PackageEvent struct {
	Action  *string            `json:"action,omitempty"`
	Repo    *github.Repository `json:"repository,omitempty"`
	Package *Package           `json:"package,omitempty"`
}
