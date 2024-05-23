package ecr

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	smTypes "github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
)

type AWSCredentials struct {
	AccessKeyID     string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`
	SessionToken    string `json:"session_token"`
}

func getAWSCredentials(secretName string) (*AWSCredentials, error) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Printf("Unable to load SDK config: %v", err)
		return nil, err
	}

	svc := secretsmanager.NewFromConfig(cfg)

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}

	result, err := svc.GetSecretValue(context.Background(), input)
	if err != nil {
		var resourceNotFoundException *smTypes.ResourceNotFoundException
		if ok := errorAs(err, &resourceNotFoundException); ok {
			log.Printf("The requested secret %s was not found", secretName)
			return nil, err
		}

		log.Printf("Failed to retrieve secret: %v", err)
		return nil, err
	}

	var credentials AWSCredentials
	err = json.Unmarshal([]byte(*result.SecretString), &credentials)
	if err != nil {
		log.Printf("Failed to unmarshal secret: %v", err)
		return nil, err
	}

	return &credentials, nil
}

func errorAs(err error, target interface{}) bool {
	if target == nil {
		return false
	}
	return errors.As(err, target)
}

func CreateRepo(repoName string) error {
	secretName := "gotask1"

	_, err := getAWSCredentials(secretName)
	if err != nil {
		log.Printf("Failed to get AWS credentials: %v", err)
		return err
	}

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Printf("Unable to load SDK config: %v", err)
		return err
	}

	svc := ecr.NewFromConfig(cfg)

	input := &ecr.CreateRepositoryInput{
		RepositoryName:     aws.String(repoName),
		ImageTagMutability: types.ImageTagMutabilityImmutable,
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
