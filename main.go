package main

import (
	"multilingo/application"
	"multilingo/entity/config"
	"multilingo/logger"
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
