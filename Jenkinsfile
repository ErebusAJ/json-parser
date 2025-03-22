@Library("git-shared") _
pipeline{
    agent { label "bub"}

    stages{
        stage("Code"){
            steps{
                script{
                    clone("https://github.com/ErebusAJ/json-parser/", "master")
                }
            }
        }

        stage("Build"){
            steps{
                script{
                    docker-build("json-parser", "latest")
                }
            }
        }

        stage("Push"){
            steps{
                script{
                    docker-push("dockerCreds", "json-parser", "latest")
                }
            }
        }

        stage("Deploy"){
            steps{
                script{
                    docker-deploy()
                }
            }
        }
    }
}