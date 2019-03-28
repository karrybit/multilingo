package main

import (
	"github.com/TakumiKaribe/multilingo/application"
	"github.com/TakumiKaribe/multilingo/entity/config"
	"github.com/TakumiKaribe/multilingo/logger"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sirupsen/logrus"
)

func main() {
	logger.Log.SetFormatter(&logrus.JSONFormatter{})
	err := config.Load()
	if err != nil {
		logger.Log.Warn(err.Error())
	}

	if config.SharedConfig.Debug {
		application.ExecDebug()

	} else {
		lambda.Start(application.LambdaHandler)
	}
}
