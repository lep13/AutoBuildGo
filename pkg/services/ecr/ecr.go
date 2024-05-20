package ecr

import (
    "context"
    "fmt"
    "log"

    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/ecr"
)


    // CreateRepository creates a new ECR repository
func CreateRepository(region, repoName string) {
    // Load the AWS configuration with the provided region
    cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
    if err != nil {
        log.Fatalf("unable to load SDK config, %v", err)
    }

    // Create a new ECR client with the loaded configuration
    svc := ecr.NewFromConfig(cfg)

    // Create the ECR repository
    input := &ecr.CreateRepositoryInput{
        RepositoryName: aws.String(repoName),
    }

    result, err := svc.CreateRepository(context.TODO(), input)
    if err != nil {
        log.Fatalf("failed to create repository, %v", err)
    }

    fmt.Printf("Repository %s created with URI: %s\n", *result.Repository.RepositoryName, *result.Repository.RepositoryUri)
}


