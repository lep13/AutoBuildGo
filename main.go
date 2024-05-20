package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/lep13/AutoBuildGo/cmd/api"
	"github.com/lep13/AutoBuildGo/ecr"
)

const (
	host = "localhost"
	port = "8080"
)

func main() {

	createRepo := flag.Bool("create-repo", false, "Create ECR repository")
	region := flag.String("region", "ap-south-1", "AWS region")
	repoName := flag.String("repo-name", "ecr-repo-1", "ECR repository name")

	flag.Parse()

	if *createRepo {
		if *repoName == "" {
			log.Fatal("repository name is required")
		}
		ecr.CreateRepository(*region, *repoName)
		return
	}

	log.Println("No action specified. Use -create-repo flag to create ECR repository.")

	err := api.ServeHTTP(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatalf("[ERROR]: %s", err)
	}

}
