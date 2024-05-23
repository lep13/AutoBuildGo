package ecr

import (
	"testing"
)

func TestCreateRepo(t *testing.T) {
	tests := []struct {
		name          string
		repoName      string
		expectedError bool
	}{
		{
			name:          "Valid Repo Name",
			repoName:      "valid-repo",
			expectedError: false,
		},
		{
			name:          "Repository Already Exists",
			repoName:      "valid-repo",
			expectedError: false, // This scenario should not return an error
		},
		{
			name:          "Invalid Repo Name: Single Character",
			repoName:      "r",
			expectedError: true,
		},
		{
			name:          "Invalid Repo Name: Using Uppercase Letters",
			repoName:      "VALID-REPO-NAME",
			expectedError: true,
		},
		{
			name:          "Valid Repo Name: Contains Underscore",
			repoName:      "valid_repo_name",
			expectedError: false,
		},
		{
			name:          "Valid Repo Name: Contains Hyphen",
			repoName:      "valid-repo-name",
			expectedError: false,
		},
		{
			name:          "Valid Repo Name: Contains Digits",
			repoName:      "repo123",
			expectedError: false,
		},
		{
			name:          "Invalid Repo Name: Empty",
			repoName:      "",
			expectedError: true,
		},
		{
			name:          "Invalid Repo Name: Contains Special Characters",
			repoName:      "repo@name",
			expectedError: true,
		},
		{
			name:          "Invalid Repo Name: Uppercase Letters",
			repoName:      "InvalidRepoName",
			expectedError: true,
		},
		{
			name:          "Invalid Repo Name: Leading Hyphen",
			repoName:      "-invalid-repo-name",
			expectedError: true,
		},
		{
			name:          "Invalid Repo Name: Trailing Hyphen",
			repoName:      "invalid-repo-name-",
			expectedError: true,
		},
		{
			name:          "Invalid Repo Name: Multiple Consecutive Hyphens",
			repoName:      "invalid--repo--name",
			expectedError: true,
		},
		{
			name:          "Invalid Repo Name: Leading Slash",
			repoName:      "/invalid-repo-name",
			expectedError: true,
		},
		{
			name:          "Invalid Repo Name: Trailing Slash",
			repoName:      "invalid-repo-name/",
			expectedError: true,
		},
		{
			name:          "Invalid Repo Name: Contains Spaces",
			repoName:      "invalid repo name",
			expectedError: true,
		},
		{
			name:          "Invalid Repo Name: Contains Non-ASCII Characters",
			repoName:      "répö-nämé",
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CreateRepo(tt.repoName)

			if (err != nil) != tt.expectedError {
				t.Errorf("CreateRepo() error = %v, expectedError %v", err, tt.expectedError)
			}
		})
	}
}
