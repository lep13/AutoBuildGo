package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/lep13/AutoBuildGo/cmd/api"
	"github.com/lep13/AutoBuildGo/pkg/services/ecr"
)

const (
	host = "localhost"
	port = "8080"
)

func main() {
	createRepo := flag.Bool("create-repo", false, "Create ECR repository")
	pushImage := flag.Bool("push-image", false, "Push Docker image to ECR repository")
	region := flag.String("region", "ap-south-1", "AWS region")
	repoName := flag.String("repo-name", "ecr-repo-1", "ECR repository name")
	imageTag := flag.String("image-tag", "", "Docker image tag")

	flag.Parse()

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

	log.Println("No action specified. Use -create-repo or -push-image flag to perform the desired action.")

	err := api.ServeHTTP(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatalf("[ERROR]: %s", err)
	}
}
