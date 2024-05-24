package ecr

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	smTypes "github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
)

type AWSCredentials struct {
	AccessKeyID     string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`
	SessionToken    string `json:"session_token"`
}

func GetAWSCredentials(secretName string) (*AWSCredentials, error) {
	cfg, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	svc := GetSecretsManagerClient(cfg)

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

	// Parse the secret assuming it is a map
	var secretMap map[string]string
	err = json.Unmarshal([]byte(*result.SecretString), &secretMap)
	if err != nil {
		log.Printf("Failed to unmarshal secret: %v", err)
		return nil, err
	}

	// Convert the map to AWSCredentials
	var credentials AWSCredentials
	for k, v := range secretMap {
		credentials.AccessKeyID = k
		credentials.SecretAccessKey = v
		// Assume no session token in this case
		credentials.SessionToken = ""
		break
	}

	return &credentials, nil
}