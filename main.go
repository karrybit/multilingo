package main

import (
	"github.com/TakumiKaribe/multilingo/application"
	"github.com/TakumiKaribe/multilingo/entity/config"
	"github.com/aws/aws-lambda-go/lambda"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := config.Load()
	if err != nil {
		log.Warn(err.Error())
	}

	if config.SharedConfig.Debug {
		application.ExecDebug()

	} else {
		lambda.Start(application.LambdaHandler)
	}
}
