pipeline {
    agent any

    parameters {
        string(name: 'AWS_REGION', defaultValue: 'your-region', description: 'AWS Region for ECR')
        string(name: 'ECR_REPO_NAME', defaultValue: 'your-ecr-repo', description: 'ECR Repository Name')
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }
        stage('Build') {
            steps {
                script {
                    // Build the Go application
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
                    // Build Docker image
                    sh "docker build -t ${params.ECR_REPO_NAME} ."
                    // Tag Docker image
                    sh "docker tag ${params.ECR_REPO_NAME} ${params.AWS_ACCOUNT_ID}.dkr.ecr.${params.AWS_REGION}.amazonaws.com/${params.ECR_REPO_NAME}:latest"
                    // Push Docker image to ECR
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
