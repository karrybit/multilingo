package main

import (
	"flag"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	if true {
		lambda.Start(LambdaHandler)
	} else {
		execDebug(flag.Arg(0), flag.Arg(1), flag.Arg(2), flag.Arg(3), flag.Arg(4))
	}
}
