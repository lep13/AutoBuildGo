package gitsetup

import (
	"testing"
)

func TestDefaultRepoConfig(t *testing.T) {
	testCases := []struct {
		name               string     // test case name
		repoName           string     // input repository name
		expectedRepoConfig RepoConfig // expected RepoConfig
	}{
		{
			name:     "Test with standard name",
			repoName: "test-repo",
			expectedRepoConfig: RepoConfig{
				Name:        "test-repo",
				Description: "Created from a template via automated setup",
				Private:     true,
				AutoInit:    true,
			},
		},
		{
			name:     "Test with empty name",
			repoName: "",
			expectedRepoConfig: RepoConfig{
				Name:        "",
				Description: "Created from a template via automated setup",
				Private:     true,
				AutoInit:    true,
			},
		},
		{
			name:     "Test with special characters",
			repoName: "repo@123",
			expectedRepoConfig: RepoConfig{
				Name:        "repo@123",
				Description: "Created from a template via automated setup",
				Private:     true,
				AutoInit:    true,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			config := DefaultRepoConfig(tc.repoName)
			if config.Name != tc.expectedRepoConfig.Name ||
				config.Description != tc.expectedRepoConfig.Description ||
				config.Private != tc.expectedRepoConfig.Private ||
				config.AutoInit != tc.expectedRepoConfig.AutoInit {
				t.Errorf("DefaultRepoConfig(%s) = %+v, want %+v", tc.repoName, config, tc.expectedRepoConfig)
			}
		})
	}
}
