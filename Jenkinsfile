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
                    dockerBuild("aaryaj", "json-parser", "latest")
                }
            }
        }

        stage("Push"){
            steps{
                script{
                    dockerPush("dockerCreds", "json-parser", "latest")
                }
            }
        }

        stage("Deploy"){
            steps{
                script{
                    dockerDeploy()
                }
            }
        }
    }
}