// config.go
package ecr

import (
	"context"
	"log"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func LoadConfig() (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Printf("Unable to load SDK config: %v", err)
		return cfg, err
	}
	return cfg, nil
}

func GetSecretsManagerClient(cfg aws.Config) *secretsmanager.Client {
	return secretsmanager.NewFromConfig(cfg)
}

func errorAs(err error, target interface{}) bool {
	if target == nil {
		return false
	}
	return errors.As(err, target)
}
