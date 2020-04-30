pipeline{
    agent any
    environment {
            DOCKER_IMAGE = "cjburchell/ridemanager"
            DOCKER_TAG = "${env.BRANCH_NAME}-${env.BUILD_NUMBER}"
            PROJECT_PATH = "/code"
    }

    parameters {
            booleanParam(name: 'UnitTests', defaultValue: false, description: 'Should unit tests run?')
    		booleanParam(name: 'Lint', defaultValue: true, description: 'Should Lint run?')
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

        stage('Static Analysis') {
                  when { expression { params.Lint } }
                  parallel {
                      stage('Vet') {
                          agent {
                              docker {
                                  image 'cjburchell/goci:1.13'
                                  args '-v $WORKSPACE:$PROJECT_PATH'
                              }
                          }
                          steps {
                              script{
                                      sh """cd ${PROJECT_PATH}/server && go list ./... | grep -v /vendor/ > projectPaths"""
                                      def paths = sh returnStdout: true, script:"""awk '{printf "/go/src/%s ",\$0} END {print ""}' projectPaths"""

                                      sh """go vet ${PROJECT_PATH}/server/..."""

                                      def checkVet = scanForIssues tool: [$class: 'GoVet']
                                      publishIssues issues:[checkVet]
                              }
                          }
                      }

                      stage('Lint') {
                          agent {
                              docker {
                                  image 'cjburchell/goci:1.13'
                                  args '-v $WORKSPACE:$PROJECT_PATH'
                              }
                          }
                          steps {
                              script{
                                  sh """cd ${PROJECT_PATH} && go list ./... | grep -v /vendor/ > projectPaths"""
                                  def paths = sh returnStdout: true, script:"""awk '{printf "/go/src/%s ",\$0} END {print ""}' projectPaths"""

                                  sh """golint ${paths}"""

                                  def checkLint = scanForIssues tool: [$class: 'GoLint']
                                  publishIssues issues:[checkLint]
                              }
                          }
                      }
                  }
        }

        stage('Tests') {
            when { expression { params.UnitTests } }
            agent {
                                          docker {
                                              image 'cjburchell/goci:1.13'
                                              args '-v $WORKSPACE:$PROJECT_PATH'
                                          }
                                      }
            steps {
                script{
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

        stage('Build') {
            steps {
                script {
                    def image = docker.build("${DOCKER_IMAGE}")
                    image.tag("${DOCKER_TAG}")
                    if( env.BRANCH_NAME == "master") {
                        image.tag("latest")
                    }
                }
            }
        }

        stage ('Push') {
            steps {
                script {
                    docker.withRegistry('', 'dockerhub') {
                       def image = docker.image("${DOCKER_IMAGE}")
                       image.push("${DOCKER_TAG}")
                       if( env.BRANCH_NAME == "master") {
                            image.push("latest")
                       }
                    }
                }
            }
        }
    }

    post {
        always {
              script{
				  sh "docker system prune -f || true"
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