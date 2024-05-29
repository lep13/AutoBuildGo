
# GoAutoSetup

## Description
GoAutoSetup is a utility tool that sets up a new software project by automatically creating a Git repository with a basic Golang service template and an Elastic Container Registry (ECR) on AWS. This enables a quick start and standardized initial setup for Golang-based projects.

## Prerequisites
Before using GoAutoSetup, ensure you have the following prerequisites configured:
- AWS Account
- Configured AWS CLI
- A secret stored in AWS Secrets Manager:
  - `github_token`: Your GitHub access token.

## Components Used
- **GitHub Repositories**: Automates the creation and setup of new repositories with standard Golang templates.
- **AWS Elastic Container Registry (ECR)**: Automates the creation of ECR for Docker container management.
- **AWS Secrets Manager**: Utilizes secrets for secure automation processes.
- **Go Modules**: Uses Go modules to manage dependencies.
- **GitHub Actions**: Implements CI/CD pipelines for automated testing and deployment.

## Usage

### Running the Application

#### Command-Line Mode:

To run the application and create a repository using the command line, use the following command:

\`\`\`sh
go run main.go <repo-name> ["optional description"]
\`\`\`

#### Web Server Mode:

To start the application as a web server, use the following command without any arguments:

\`\`\`sh
go run main.go
\`\`\`

Once the server is running, you can create a repository by making a POST request to the server's endpoint:

\`\`\`sh
curl -X POST -H "Content-Type: application/json" -d '{"repo_name": "test-repo", "description": "A test repository"}' http://localhost:8080/create-repo
\`\`\`

Ensure the repository name is in the correct format as specified:
The repository name should consist of lowercase letters and numbers, optionally separated by dots, underscores, or hyphens, and can include slashes to indicate subdirectories.

### Testing

To execute tests for the ECR and GitHub functionalities, navigate to the directory containing the respective test case files and run:

\`\`\`sh
go test ./...
\`\`\`

For more details and updates, refer to the [project's GitHub page](https://github.com/lep13/ServiceTemplate).
