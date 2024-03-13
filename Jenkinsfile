pipeline {
  agent any

  environment {
      VERSION = sh(script: 'git describe --tags --abbrev=0', returnStdout: true).trim()
  }
    
  stages {
    stage("Version"){
      steps {
        echo "$VERSION"
      }
    }
    
    stage("Docker Build") {
      steps {
        sh("docker build -t go-api:v10 .")
      }
    }

    stage("Docker Login") {
      steps {
        sh("docker login")
      }
    }

    stage("Run Docker Image") {
      steps {
        sh("docker run -d -p 80:8080 go-api:v10")
      }
    }
  }
}
