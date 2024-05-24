package ecr

import (
	"context"
	"log"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
)

func CreateRepo(repoName string) error {
	secretName := "gotask1"

	creds, err := GetAWSCredentials(secretName)
	if err != nil {
		log.Printf("Failed to get AWS credentials: %v", err)
		return err
	}

	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithCredentialsProvider(
			aws.NewCredentialsCache(
				aws.CredentialsProviderFunc(func(context.Context) (aws.Credentials, error) {
					return aws.Credentials{
						AccessKeyID:     creds.AccessKeyID,
						SecretAccessKey: creds.SecretAccessKey,
						SessionToken:    creds.SessionToken,
					}, nil
				}),
			),
		),
	)
	if err != nil {
		log.Printf("Unable to load SDK config: %v", err)
		return err
	}

	svc := ecr.NewFromConfig(cfg)

	input := &ecr.CreateRepositoryInput{
		RepositoryName:     aws.String(repoName),
		ImageTagMutability: types.ImageTagMutabilityImmutable,
		ImageScanningConfiguration: &types.ImageScanningConfiguration{
			ScanOnPush: true,
		},
	}

	_, err = svc.CreateRepository(context.Background(), input)
	if err != nil {
		var repoAlreadyExistsErr *types.RepositoryAlreadyExistsException
		if errors.As(err, &repoAlreadyExistsErr) {
			log.Printf("Repository %s already exists.", repoName)
			return nil
		}
		log.Printf("Failed to create repository: %v", err)
		return err
	}
	return nil
}
