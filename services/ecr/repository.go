package ecr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	// "strings"

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
		return nil, fmt.Errorf("unable to load SDK config, %v", err)
	}

	svc := secretsmanager.NewFromConfig(cfg)

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}

	result, err := svc.GetSecretValue(context.Background(), input)
	if err != nil {
		var resourceNotFoundException *smTypes.ResourceNotFoundException
		if ok := errorAs(err, &resourceNotFoundException); ok {
			return nil, fmt.Errorf("the requested secret %s was not found", secretName)
		}

		return nil, fmt.Errorf("failed to retrieve secret: %v", err)
	}

	var credentials AWSCredentials
	err = json.Unmarshal([]byte(*result.SecretString), &credentials)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal secret: %v", err)
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
		return fmt.Errorf("failed to get AWS credentials: %v", err)
	}

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return fmt.Errorf("unable to load SDK config, %v", err)
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
		return fmt.Errorf("failed to create repository: %v", err)
	}
	return nil
}
