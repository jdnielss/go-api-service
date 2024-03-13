pipeline {
  agent any

  stages {
    stage("Docker Build") {
      steps {
        sh("docker build -t go-api:v1 .")
      }
    }

    stage("Docker Login") {
      steps {
        sh("docker login")
      }
    }

    stage("Push Docker Image") {
      steps {
        sh("docker push go-api:v1")
      }
    }

    stage("Run Docker Image") {
      steps {
        sh("docker run -p 8080:8080 go-api:v1 -d")
      }
    }
  }
}
