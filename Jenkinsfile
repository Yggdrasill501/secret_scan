pipeline {
    agent any

    environment {
        BITBUCKET_REPO = 'your-bitbucket-repo-url'
        BITBUCKET_CREDENTIALS_ID = 'your-credentials-id'
        GITLEAKS_VERSION = 'latest'
    }

    stages {
        stage('Checkout') {
            steps {
                // Checkout the Bitbucket repository
                git url: "${env.BITBUCKET_REPO}", credentialsId: "${env.BITBUCKET_CREDENTIALS_ID}"
            }
        }

        stage('Scan for Secrets') {
            steps {
                // Install Gitleaks
                sh 'curl -sSfL https://github.com/zricethezav/gitleaks/releases/download/${GITLEAKS_VERSION}/gitleaks-linux-amd64 -o gitleaks'
                sh 'chmod +x gitleaks'
                
                // Run Gitleaks scan
                sh './gitleaks detect --source . --report-format json --report-path gitleaks_report.json'
            }
        }

        stage('Publish Results') {
            steps {
                // Archive the Gitleaks report
                archiveArtifacts artifacts: 'gitleaks_report.json', allowEmptyArchive: true
            }
        }
    }

    post {
        always {
            cleanWs()
        }
    }
}
