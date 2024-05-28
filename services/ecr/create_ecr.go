package ecr

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
)

// AWSClient interface defines the methods from ecr.Client that are used in CreateRepo function.
type AWSClient interface {
	CreateRepository(ctx context.Context, params *ecr.CreateRepositoryInput, optFns ...func(*ecr.Options)) (*ecr.CreateRepositoryOutput, error)
}

// CreateRepo creates a repository in Amazon ECR using the provided AWS client.
func CreateRepo(repoName string, client AWSClient) error {
	input := &ecr.CreateRepositoryInput{
		RepositoryName:     aws.String(repoName),
		ImageTagMutability: types.ImageTagMutabilityImmutable,
		ImageScanningConfiguration: &types.ImageScanningConfiguration{
			ScanOnPush: true,
		},
	}

	_, err := client.CreateRepository(context.Background(), input)
	if err != nil {
		var repoAlreadyExistsErr *types.RepositoryAlreadyExistsException
		if errors.As(err, &repoAlreadyExistsErr) {
			log.Printf("Repository %s already exists.", repoName)
			return nil // Return early if repository already exists
		}
		log.Printf("Failed to create repository: %v", err)
		return err
	}
	return nil
}
