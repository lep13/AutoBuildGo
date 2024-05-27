package gitsetup

// default repository configuration with a dynamic description.
func DefaultRepoConfig(repoName string, description string) RepoConfig {
    return RepoConfig{
        Name:        repoName,
        Description: description,
        Private:     true,
        AutoInit:    true,
        TemplateURL: "https://api.github.com/repos/lep13/ServiceTemplate/generate",
    }
}
