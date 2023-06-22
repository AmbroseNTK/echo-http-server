pipeline {
   agent {
        kubernetes {
            label 'docker-builder'
            yaml """
            apiVersion: v1
            kind: Pod
            metadata:
              labels:
                app: docker-builder
            spec:
              containers:
              - name: docker
                image: docker:latest
            """
        }
    }

    stages {
        stage('Build') {
            steps {
                // Build the Docker image
                sh "docker build -t echohttp:${env.BUILD_NUMBER} ."
            }
        }
        // stage('Push') {
        //     steps {
        //         // Push the Docker image to a registry
        //         withCredentials([usernamePassword(credentialsId: 'docker-hub-credentials', usernameVariable: 'DOCKER_USERNAME', passwordVariable: 'DOCKER_PASSWORD')]) {
        //             sh 'docker login -u $DOCKER_USERNAME -p $DOCKER_PASSWORD'
        //             sh "docker push ambrosentk/echohttp:${env.BUILD_NUMBER}"
        //         }
        //     }
        // }
        // stage('Deploy') {
        //     steps {
        //         // Update the Kubernetes deployment with the new image
        //         withCredentials([file(credentialsId: 'k8s-config', variable: 'KUBECONFIG')]) {
        //             sh "kubectl config use-context my-context --kubeconfig=$KUBECONFIG"
        //             sh "kubectl set image deployment/echohttp container=ambrosentk/echohttp:${env.BUILD_NUMBER}"
        //         }
        //     }
        // }
    }
}