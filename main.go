package main

import (
	"github.com/TakumiKaribe/multilingo/application"
	"github.com/TakumiKaribe/multilingo/entity/config"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = config.Load()
	if err != nil {
		log.Warn(err.Error())
	}

	if config.SharedConfig.Debug {
		application.ExecDebug()

	} else {
		lambda.Start(application.LambdaHandler)
	}
}
