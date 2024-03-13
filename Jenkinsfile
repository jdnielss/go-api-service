pipeline {
  agent any

  stages {
    stage("Docker Build") {
      sh("docker build -t go-api:v1 .")
    }

    stage("Push Docker Image") {
      sh("docker push go-api:v1")
    }

    stage("Run Docker Image") {
      sh("docker run -p 8080:8080 go-api:v1 -d")
    }
  }
}
