node ('master') {
    stage ('Checkout') {
        checkout scm
    }
    stage ('Build') {
        sh """
            gcloud container builds submit --tag gcr.io/matt-sandbox-155112/blue-green:${env.BUILD_NUMBER} .
        """
    }
}
