package ecr

import (
	"testing"
)

func TestCreateRepository(t *testing.T) {
	type args struct {
		region   string
		repoName string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ValidCase",
			args: args{
				region:   "us-west-2",
				repoName: "test-repo",
			},
		},
		{
			name: "EmptyRegion",
			args: args{
				region:   "",
				repoName: "test-repo",
			},
		},
		{
			name: "EmptyRepoName",
			args: args{
				region:   "us-west-2",
				repoName: "",
			},
		},
		// Add more test cases for different scenarios
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateRepository(tt.args.region, tt.args.repoName)
			// Add assertions here if necessary
		})
	}
}

func TestPushImageToRepository(t *testing.T) {
	type args struct {
		region   string
		repoName string
		imageTag string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ValidPush",
			args: args{
				region:   "us-west-2",
				repoName: "test-repo",
				imageTag: "test-image:latest",
			},
			wantErr: false,
		},
		{
			name: "EmptyRegion",
			args: args{
				region:   "",
				repoName: "test-repo",
				imageTag: "test-image:latest",
			},
			wantErr: true,
		},
		{
			name: "EmptyRepoName",
			args: args{
				region:   "us-west-2",
				repoName: "",
				imageTag: "test-image:latest",
			},
			wantErr: true,
		},
		{
			name: "EmptyImageTag",
			args: args{
				region:   "us-west-2",
				repoName: "test-repo",
				imageTag: "",
			},
			wantErr: true,
		},
		// Add more test cases for failure scenarios
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PushImageToRepository(tt.args.region, tt.args.repoName, tt.args.imageTag); (err != nil) != tt.wantErr {
				t.Errorf("PushImageToRepository() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetAuthorizationToken(t *testing.T) {
	type args struct {
		region string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ValidCase",
			args: args{
				region: "us-west-2",
			},
			wantErr: false,
		},
		{
			name: "EmptyRegion",
			args: args{
				region: "",
			},
			wantErr: true,
		},
		// Add more test cases for failure scenarios
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := GetAuthorizationToken(tt.args.region)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAuthorizationToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

