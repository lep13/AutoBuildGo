package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/lep13/AutoBuildGo/pkg/services/Repo"
	"github.com/lep13/AutoBuildGo/pkg/services/ecr"
)

const (
	host = "localhost"
	port = "8080"
)

func main() {
	token := flag.String("token", "", "GitHub personal access token")
	gitrepoName := flag.String("git-repo-name", "", "Name of the GitHub repository")
	description := flag.String("description", "", "Description of the GitHub repository")
	isPrivate := flag.Bool("private", false, "Make the GitHub repository private")

	createRepo := flag.Bool("create-repo", false, "Create ECR repository")
	pushImage := flag.Bool("push-image", false, "Push Docker image to ECR repository")
	region := flag.String("region", "ap-south-1", "AWS region")
	repoName := flag.String("repo-name", "ecr-repo-1", "ECR repository name")
	imageTag := flag.String("image-tag", "latest", "Docker image tag")

	flag.Parse()
	if *token != "" && *gitrepoName != "" && *description != "" {
		// Call the function to create the GitHub repository
		err := Repo.CreateGitHubRepo(*token, *gitrepoName, *description, *isPrivate)
		if err != nil {
			log.Fatal("Error creating GitHub repository:", err)
		}

		fmt.Println("GitHub repository created successfully!")
	}

	if *createRepo {
		if *repoName == "" {
			log.Fatal("repository name is required")
		}
		ecr.CreateRepository(*region, *repoName)
		return
	}

	if *pushImage {
		if *repoName == "" || *imageTag == "" {
			log.Fatal("both repository name and image tag are required")
		}
		err := ecr.PushImageToRepository(*region, *repoName, *imageTag)
		if err != nil {
			log.Fatalf("failed to push image to repository: %v", err)
		}
		return
	}

	// log.Println("No action specified. Use -create-repo or -push-image  flag to perform the desired action.")

	// err1 := api.ServeHTTP(fmt.Sprintf("%s:%s", host, port))
	// if err1 != nil {
	// 	log.Fatalf("[ERROR]: %s", err1)
	// }
}
