package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	log "github.com/sirupsen/logrus"

	"github.com/TakumiKaribe/multilingo/request/slack"
)

func main() {
	if true {
		lambda.Start(LambdaHandler)
	} else {
		slackClient, err := slack.NewClient("host", "token")
		if err != err {
			log.Fatal(err.Error())
		}

		body := slack.SlackRequestBody{}
		body.Token = "token"
		attachment := slack.Attachment{}
		attachment.Color = "good"
		attachment.Title = "Dummy Title"
		attachment.TitleLink = "https://github.com/TakumiKaribe/multilingo"
		attachment.Text = "result.Response.BuildStdout"
		body.Attachments = append(body.Attachments, &attachment)
		body.Channel = "channel"
		body.UserName = "requestBody.Event.User"

		slackClient.Notification(body)
	}
}
