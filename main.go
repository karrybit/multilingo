package main

import (
	"github.com/TakumiKaribe/multilingo/running"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	if false {
		lambda.Start(running.LambdaHandler)
	} else {
		running.ExecDebug()
	}
}
