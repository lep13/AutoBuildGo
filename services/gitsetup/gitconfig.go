package gitsetup

// default configuration for repository creation.
func DefaultRepoConfig(repoName string) RepoConfig {
	return RepoConfig{
		Name:        repoName,
		Description: "Created from a template via automated setup",
		Private:     true,
		AutoInit:    true,
	}
}
