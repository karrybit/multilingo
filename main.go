package main

import (
	"os"
	"strconv"

	"github.com/TakumiKaribe/multilingo/application"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if b, _ := strconv.ParseBool(os.Getenv("DEBUG")); b {
		application.ExecDebug()

	} else {
		lambda.Start(application.LambdaHandler)
	}
}
