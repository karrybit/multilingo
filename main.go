package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	if false {
		lambda.Start(LambdaHandler)
	} else {
		execDebug()
	}
}
