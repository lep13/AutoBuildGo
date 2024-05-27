package gitsetup

type RepoConfig struct {
	Name        string
	Description string
	Private     bool
	AutoInit    bool
	TemplateURL string
}

type SecretData struct {
	GITHUB_TOKEN string `json:"GITHUB_TOKEN"`
}
