package ecr

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	ecrpublic "github.com/aws/aws-sdk-go-v2/service/ecrpublic"
	//  "github.com/aws/aws-sdk-go-v2/service/ecrpublic/types"
)

// CreatePublicRepository creates a new public ECR repository
func CreatePublicRepository(region, repoName string) {
	// Load the AWS configuration with the provided region
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Create a new ECR public client with the loaded configuration
	svc := ecrpublic.NewFromConfig(cfg)

	// Create the ECR public repository
	input := &ecrpublic.CreateRepositoryInput{
		RepositoryName: aws.String(repoName),
	}

	result, err := svc.CreateRepository(context.TODO(), input)
	if err != nil {
		log.Fatalf("failed to create repository, %v", err)
	}

	repoURI := *result.Repository.RepositoryUri
	fmt.Printf("Public repository %s created with URI: %s\n", *result.Repository.RepositoryName, repoURI)
}

// PushImageToRepository pushes a Docker image to the ECR repository
func PushImageToRepository(region, repoName, imageTag string) error {
	// Get the ECR authorization token
	authToken, proxyEndpoint, err := GetAuthorizationToken(region)
	if err != nil {
		return fmt.Errorf("failed to get authorization token: %w", err)
	}

	// Decode the authorization token (assuming base64 encoded)
	decodedToken, err := base64.StdEncoding.DecodeString(authToken)
	if err != nil {
		return fmt.Errorf("failed to decode authorization token: %w", err)
	}

	// Extract the username and password from the authorization token
	tokenParts := strings.SplitN(string(decodedToken), ":", 2)
	if len(tokenParts) != 2 {
		return fmt.Errorf("invalid authorization token format")
	}
	username := tokenParts[0]
	password := tokenParts[1]

	// Construct the image URI without the https:// prefix
	repoURI := strings.TrimPrefix(proxyEndpoint, "https://")
	imageURI := fmt.Sprintf("%s/%s:%s", repoURI, repoName, imageTag)

	// Configure Docker to use the ECR registry
	cmd := exec.Command("docker", "login", "-u", username, "-p", password, proxyEndpoint)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to authenticate Docker to ECR: %s, %w", string(output), err)
	}

	// Tag the Docker image
	fmt.Println("Tagging image", imageTag, "to URI", imageURI)
	cmd = exec.Command("docker", "tag", imageTag, imageURI)
	var tagOutput bytes.Buffer
	cmd.Stdout = &tagOutput
	cmd.Stderr = &tagOutput
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to tag Docker image: %s, %w", tagOutput.String(), err)
	}

	// Push the Docker image
	cmd = exec.Command("docker", "push", imageURI)
	var pushOutput bytes.Buffer
	cmd.Stdout = &pushOutput
	cmd.Stderr = &pushOutput
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to push Docker image: %s, %w", pushOutput.String(), err)
	}

	fmt.Printf("Image %s pushed to repository %s\n", imageTag, imageURI)
	return nil
}

func GetAuthorizationToken(region string) (string, string, error) {
	// Load the AWS configuration with the provided region
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return "", "", fmt.Errorf("unable to load SDK config: %w", err)
	}

	// Create a new ECR client with the loaded configuration
	svc := ecr.NewFromConfig(cfg)

	// Get the ECR authorization token
	input := &ecr.GetAuthorizationTokenInput{}
	result, err := svc.GetAuthorizationToken(context.TODO(), input)
	if err != nil {
		return "", "", fmt.Errorf("failed to get authorization token: %w", err)
	}
	if len(result.AuthorizationData) == 0 {
		return "", "", fmt.Errorf("no authorization data found")
	}

	authData := result.AuthorizationData[0]
	authToken := aws.ToString(authData.AuthorizationToken)
	proxyEndpoint := aws.ToString(authData.ProxyEndpoint)

	return authToken, proxyEndpoint, nil
}
