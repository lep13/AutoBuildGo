// config.go

package ecr

import (
	"context"
	"log"

	// "github.com/aws/aws-sdk-go-v2/config"
)

// ConfigLoader defines the interface for loading AWS SDK config
type ConfigLoader interface {
	LoadDefaultConfig(ctx context.Context) (interface{}, error)
}

// LoadConfig loads the AWS SDK config
func LoadConfig(loader ConfigLoader) (interface{}, error) {
	cfg, err := loader.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Printf("Unable to load SDK config: %v", err)
		return nil, err
	}
	return cfg, nil
}
