package ecr

import "testing"

func TestCreateRepo(t *testing.T) {
	type args struct {
		repoName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid Repo Name",
			args: args{
				repoName: "valid-repo",
			},
			wantErr: false,
		},
		{
			name: "Repository Already Exists",
			args: args{
				repoName: "valid-repo",
			},
			wantErr: false, // This scenario should not return an error
		},
		{
			name: "Invalid Repo Name: Single Character",
			args: args{
				repoName: "r",
			},
			wantErr: true,
		},
		{
			name: "Invalid Repo Name: Using Uppercase Letters",
			args: args{
				repoName: "VALID-REPO-NAME",
			},
			wantErr: true,
		},
		{
			name: "Valid Repo Name: Contains Underscore",
			args: args{
				repoName: "valid_repo_name",
			},
			wantErr: false,
		},
		{
			name: "Valid Repo Name: Contains Hyphen",
			args: args{
				repoName: "valid-repo-name",
			},
			wantErr: false,
		},
		{
			name: "Valid Repo Name: Contains Digits",
			args: args{
				repoName: "repo123",
			},
			wantErr: false,
		},
		{
			name: "Invalid Repo Name: Empty",
			args: args{
				repoName: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid Repo Name: Contains Special Characters",
			args: args{
				repoName: "repo@name",
			},
			wantErr: true,
		},
		{
			name: "Invalid Repo Name: Uppercase Letters",
			args: args{
				repoName: "InvalidRepoName",
			},
			wantErr: true,
		},
		{
			name: "Invalid Repo Name: Leading Hyphen",
			args: args{
				repoName: "-invalid-repo-name",
			},
			wantErr: true,
		},
		{
			name: "Invalid Repo Name: Trailing Hyphen",
			args: args{
				repoName: "invalid-repo-name-",
			},
			wantErr: true,
		},
		{
			name: "Invalid Repo Name: Multiple Consecutive Hyphens",
			args: args{
				repoName: "invalid--repo--name",
			},
			wantErr: true,
		},
		{
			name: "Invalid Repo Name: Leading Slash",
			args: args{
				repoName: "/invalid-repo-name",
			},
			wantErr: true,
		},
		{
			name: "Invalid Repo Name: Trailing Slash",
			args: args{
				repoName: "invalid-repo-name/",
			},
			wantErr: true,
		},
		{
			name: "Invalid Repo Name: Contains Spaces",
			args: args{
				repoName: "invalid repo name",
			},
			wantErr: true,
		},
		{
			name: "Invalid Repo Name: Contains Non-ASCII Characters",
			args: args{
				repoName: "répö-nämé",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateRepo(tt.args.repoName); (err != nil) != tt.wantErr {
				t.Errorf("CreateRepo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
