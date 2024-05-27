package ecr

import (
	"context"
	"log"


	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func LoadConfig() (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Printf("Unable to load SDK config: %v", err)
		return cfg, err
	}
	return cfg, nil
}
