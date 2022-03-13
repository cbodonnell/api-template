pipeline {
    agent any
    environment {
        GOROOT = "${tool type: 'go', name: 'go1.15.6'}/go"
    }
    stages {
        stage('build') {
            steps {
                echo 'building...'
                sh 'echo $GOROOT'
                sh '$GOROOT/bin/go build'
            }
        }
        stage('test') {
            steps {
                echo 'testing...'
            }
        }
        stage('deploy') {
            steps {
                echo 'deploying...'
                sh 'sudo systemctl stop api-template'
                sh 'sudo cp api-template /etc/api-template/api-template'
                sh 'sudo systemctl start api-template'
            }
        }
    }
    post {
        cleanup {
            deleteDir()
        }
    }
}