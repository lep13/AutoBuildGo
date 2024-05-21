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
                checkout scm
            }
        }
        stage('Setup GitHub Repo') {
            steps {
                script {
                    
                    sh "echo 'Create GitHub Repo with name: ${params.GITHUB_REPO}'"
                    // GitHub CLI or API to create a repo
                }
            }
        }
        stage('Add Templates') {
            steps {
                script {
                    // Jenkinsfile and Go project template
                    sh "echo 'Add template Jenkinsfile and Go project to ${params.GITHUB_REPO}'"
                    
                }
            }
        }
        stage('Build') {
            steps {
                script {
                    sh 'go build -o main ./main.go'
                }
            }
        }
        stage('Create ECR Repository') {
            steps {
                script {
                    sh "go run main.go --create-repo --repo-name=${params.ECR_REPO_NAME} --region=${params.AWS_REGION}"
                }
            }
        }
        stage('Docker Build and Push') {
            steps {
                script {
                    sh "docker build -t ${params.ECR_REPO_NAME} ."
                    sh "docker tag ${params.ECR_REPO_NAME} ${params.AWS_ACCOUNT_ID}.dkr.ecr.${params.AWS_REGION}.amazonaws.com/${params.ECR_REPO_NAME}:latest"
                    sh "docker push ${params.AWS_ACCOUNT_ID}.dkr.ecr.${params.AWS_REGION}.amazonaws.com/${params.ECR_REPO_NAME}:latest"
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
