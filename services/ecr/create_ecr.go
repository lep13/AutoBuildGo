package ecr

import (
    "context"
    "log"

    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/service/ecr"
    "github.com/aws/aws-sdk-go-v2/service/ecr/types"
)

type ECRClientInterface interface {
    CreateRepository(ctx context.Context, params *ecr.CreateRepositoryInput, optFns ...func(*ecr.Options)) (*ecr.CreateRepositoryOutput, error)
}

// CreateRepo creates a repository in Amazon ECR using the provided ECR client.
func CreateRepo(repoName string, ecrClient ECRClientInterface) error {
    input := &ecr.CreateRepositoryInput{
        RepositoryName:     aws.String(repoName),
        ImageTagMutability: types.ImageTagMutabilityImmutable,
        ImageScanningConfiguration: &types.ImageScanningConfiguration{
            ScanOnPush: true,
        },
    }

    _, err := ecrClient.CreateRepository(context.Background(), input)
    if err != nil {
        log.Printf("Failed to create repository: %v", err)
        return err
    }

    log.Printf("Repository %s created successfully.", repoName)
    return nil
}
