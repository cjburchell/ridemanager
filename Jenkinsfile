pipeline{
    agent any
    environment {
            DOCKER_IMAGE = "cjburchell/ridemanager"
            DOCKER_TAG = "${env.BRANCH_NAME}"
            PROJECT_PATH = "/code"
    }

    parameters {
            booleanParam(name: 'UnitTests', defaultValue: false, description: 'Should unit tests run?')
    		booleanParam(name: 'Lint', defaultValue: false, description: 'Should Lint run?')
        }

    stages{
        stage('Setup') {
            steps {
                script{
                    slackSend color: "good", message: "Job: ${env.JOB_NAME} with build number ${env.BUILD_NUMBER} started"
                }
             /* Let's make sure we have the repository cloned to our workspace */
             checkout scm
             }
         }

        stage('Lint') {
            when { expression { params.Lint } }
            steps {
                script{
                docker.withRegistry('https://390282485276.dkr.ecr.us-east-1.amazonaws.com', 'ecr:us-east-1:redpoint-ecr-credentials') {
                        docker.image('cjburchell/goci:latest').inside("-v ${env.WORKSPACE}:${PROJECT_PATH}"){
                            sh """cd ${PROJECT_PATH} && go list ./... | grep -v /vendor/ > projectPaths"""
                            def paths = sh returnStdout: true, script:"""awk '{printf "/go/src/%s ",\$0} END {print ""}' projectPaths"""

                            sh """go tool vet ${paths}"""
                            sh """golint ${paths}"""

                            warnings canComputeNew: true, canResolveRelativePaths: true, categoriesPattern: '', consoleParsers: [[parserName: 'Go Vet'], [parserName: 'Go Lint']], defaultEncoding: '', excludePattern: '', healthy: '', includePattern: '', messagesPattern: '', unHealthy: ''
                        }
                    }
                }
            }
        }

        stage('Tests') {
            when { expression { params.UnitTests } }
            steps {
                script{
                    docker.withRegistry('https://390282485276.dkr.ecr.us-east-1.amazonaws.com', 'ecr:us-east-1:redpoint-ecr-credentials') {
                        docker.image('cjburchell/goci:latest').inside("-v ${env.WORKSPACE}:${PROJECT_PATH}"){
                            sh """cd ${PROJECT_PATH} && go list ./... | grep -v /vendor/ > projectPaths"""
                            def paths = sh returnStdout: true, script:"""awk '{printf "/go/src/%s ",\$0} END {print ""}' projectPaths"""

                            def testResults = sh returnStdout: true, script:"""go test -v ${paths}"""
                            writeFile file: 'test_results.txt', text: testResults
                            sh """go2xunit -input test_results.txt > tests.xml"""
                            sh """cd ${PROJECT_PATH} && ls"""

                            archiveArtifacts 'test_results.txt'
                            archiveArtifacts 'tests.xml'
                            junit allowEmptyResults: true, testResults: 'tests.xml'
                        }
                    }
                }
            }
        }

        stage('Build') {
            steps {
                script {
                    if( env.BRANCH_NAME == "master") {
                        docker.build("${DOCKER_IMAGE}").tag("latest")
                    }
                    else {
                        docker.build("${DOCKER_IMAGE}").tag("${DOCKER_TAG}")
                    }
                }
            }
        }

        stage ('Push') {
            steps {
                script {
                    docker.withRegistry('https://390282485276.dkr.ecr.us-east-1.amazonaws.com', 'ecr:us-east-1:redpoint-ecr-credentials') {
                        if( env.BRANCH_NAME == "master")
                        {
                            docker.image("${DOCKER_IMAGE}").push("latest")
                        }
                        else {
                            docker.image("${DOCKER_IMAGE}").push("${DOCKER_TAG}")
                        }
                    }
                }
            }
        }
        stage ('Deploy') {
            agent {
                docker {
                image 'fabn/rancher-cli'
                args '--env RANCHER_URL=${RANCHER_URL} --env RANCHER_ACCESS_KEY=${RANCHER_ACCESS_KEY} --env RANCHER_SECRET_KEY=${RANCHER_SECRET_KEY}'
                 }
            }
            when {
                branch 'master'
            }
            steps {
                script {
                       sh """rancher --debug up -d --force-upgrade --pull --confirm-upgrade --stack Redpoint"""
                }
            }
        }
    }

    post {
        always {
              script{
                  if ( currentBuild.currentResult == "SUCCESS" ) {
                    slackSend color: "good", message: "Job: ${env.JOB_NAME} with build number ${env.BUILD_NUMBER} was successful"
                  }
                  else if( currentBuild.currentResult == "FAILURE" ) {
                    slackSend color: "danger", message: "Job: ${env.JOB_NAME} with build number ${env.BUILD_NUMBER} was failed"
                  }
                  else if( currentBuild.currentResult == "UNSTABLE" ) {
                    slackSend color: "warning", message: "Job: ${env.JOB_NAME} with build number ${env.BUILD_NUMBER} was unstable"
                  }
                  else {
                    slackSend color: "danger", message: "Job: ${env.JOB_NAME} with build number ${env.BUILD_NUMBER} its result (${currentBuild.currentResult}) was unclear"
                  }
              }
        }
    }

}