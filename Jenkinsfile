pipeline {
    agent any

    stages {
        stage('Unit') {
            steps {
                echo 'Unit Test..'
                sh 'go test -v'
            }
        }
        // stage('Test') {
        //     steps {
        //         echo 'Testing..'
        //     }
        // }
        // stage('Deploy') {
        //     steps {
        //         echo 'Deploying....'
        //     }
        // }
    }
}