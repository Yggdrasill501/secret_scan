# Jenkins Pipeline with Gitleaks and Go

This project contains a Jenkins pipeline defined in a Jenkinsfile and a Go script to trigger the pipeline. The pipeline checks out a Bitbucket repository, scans for secrets using Gitleaks, and archives the results.

## Jenkins Token 
```
export JENKINS_USER=your-jenkins-username
export JENKINS_API_TOKEN=your-jenkins-api-token
```
