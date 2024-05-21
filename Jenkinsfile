pipeline {
    agent any

    parameters {
        string(name: 'AWS_REGION', defaultValue: 'your-region', description: 'AWS Region for ECR')
        string(name: 'ECR_REPO_NAME', defaultValue: 'your-ecr-repo', description: 'ECR Repository Name')
        string(name: 'AWS_ACCOUNT_ID', defaultValue: 'your-account-id', description: 'AWS Account ID')
        string(name: 'GITHUB_REPO', defaultValue: 'new-repo', description: 'GitHub Repository Name')
    }

    stages {
        stage('Checkout') {
            steps {
                script {
                    bat 'git clone https://github.com/lep13/AutoBuildGo || exit 1'
                }
            }
        }
        stage('Setup GitHub Repo') {
            steps {
                script {
                    // Create GitHub Repo using GitHub CLI
                    bat "gh repo create new-repo --public --confirm"
                }
            }
        }
        stage('Add Templates') {
            steps {
                script {
                    // Clone and pubat Jenkinsfile and Go project template
                    bat "git clone https://github.com/lep13/AutoBuildGo"
                    bat "cd AutoBuildGo && git remote set-url origin https://github.com/${params.GITHUB_REPO}"
                    bat "cd AutoBuildGo && git pubat -u origin master"
                }
            }
        }
        stage('Build') {
            steps {
                script {
                    bat 'go build -o main ./main.go'
                }
            }
        }
        stage('Create ECR Repository') {
            steps {
                script {
                    bat "go run main.go --create-repo --repo-name=${params.ECR_REPO_NAME} --region=${params.AWS_REGION}"
                }
            }
        }
        stage('Docker Build and Pubat') {
            steps {
                script {
                    bat "docker build -t ${params.ECR_REPO_NAME} ."
                    bat "docker tag ${params.ECR_REPO_NAME} ${params.AWS_ACCOUNT_ID}.dkr.ecr.${params.AWS_REGION}.amazonaws.com/${params.ECR_REPO_NAME}:latest"
                    bat "docker pubat ${params.AWS_ACCOUNT_ID}.dkr.ecr.${params.AWS_REGION}.amazonaws.com/${params.ECR_REPO_NAME}:latest"
                }
            }
        }
    }
    post {
        always {
            cleanWs()
        }
    }
}
