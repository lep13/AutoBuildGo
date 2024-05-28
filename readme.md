
# AutoBuildGo

## Description
AutoBuildGo is a utility tool that setups a new software project by automatically creating a Git repository with a basic Golang service template, and an Elastic Container Registry (ECR) on AWS. This enables quick start-up and standardized initial setup for Golang-based projects.

## Prerequisites
Before using AutoBuildGo, ensure you have the following prerequisites configured:
- AWS Account
- Configured AWS CLI
- Two secrets stored in AWS Secrets Manager:
  - `github_token`: Your GitHub access token.
  - `gotask1`: Your AWS Access keys.

## Components Used
- **GitHub Repositories**: Automates the creation and setup of new repositories with standard Golang templates.
- **AWS Elastic Container Registry (ECR)**: Automates the creation of ECR for Docker container management.
- **AWS Secrets Manager**: Utilizes secrets for secure automation processes.
- **Go Modules**: Uses Go modules to manage dependencies.
- **GitHub Actions**: Implements CI/CD pipelines for automated testing and deployment.

## Commands
### Create Repositories

You can run your main.go with a Dynamic description like this:
```bash
go run main.go <repo-name> ["optional description"]

```

If you do not want to add a description, run this to create a Git Repository with the default description: 
```bash
go run main.go <reponame>

```

Ensure the repository name is in the correct format as specified:

The repository name should consist of lowercase letters and numbers, optionally separated by dots, underscores, or hyphens, and can include slashes to indicate subdirectories.


### Test
To execute tests, navigate to the directory containing the respective test case files for ECR and GitHub:
```bash
go test
```

For more details and updates, refer to the [project's GitHub page](https://github.com/lep13/ServiceTemplate).
