package main

import (
	"github.com/TakumiKaribe/multilingo/application"
	"github.com/TakumiKaribe/multilingo/entity/config"
	"github.com/TakumiKaribe/multilingo/logger"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	err := config.Load()
	if err != nil {
		logger.Log.Warn(err.Error())
		return
	}

	if config.SharedConfig.Debug {
		application.ExecDebug()

	} else {
		lambda.Start(application.LambdaHandler)
	}
}
