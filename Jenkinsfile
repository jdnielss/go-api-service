pipeline {
    agent any
    
    environment {
        GIT_TAG = sh(script: 'git describe --tags $(git rev-list --tags --max-count=1)', returnStdout: true).trim()
        DOCKER_IMAGE = "go-api:${GIT_TAG}"
    }
    
    stages {
        stage("Docker Build") {
            steps {
                sh "docker build -t ${DOCKER_IMAGE} ."
            }
        }

        stage("Docker Login") {
            steps {
                sh("docker login")
            }
        }

        stage("Run Docker Image") {
            steps {
                sh("docker run -d -p 80:8080 ${DOCKER_IMAGE}")
            }
        }
    }
}
