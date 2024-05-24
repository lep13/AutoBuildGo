package gitsetup

type RepoConfig struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
	AutoInit    bool   `json:"auto_init"`
}

type SecretData struct {
	GITHUB_TOKEN string `json:"GITHUB_TOKEN"`
}
