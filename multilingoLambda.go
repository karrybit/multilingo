package main

import (
	"context"

	"github.com/TakumiKaribe/multilingo/model"
	"github.com/TakumiKaribe/multilingo/request/paiza"
	"github.com/TakumiKaribe/multilingo/request/slack"
	"github.com/aws/aws-lambda-go/events"
	log "github.com/sirupsen/logrus"
)

// LambdaHandler -
func LambdaHandler(ctx context.Context, apiRequest events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if val, ok := apiRequest.Headers["X-Slack-Retry-Num"]; val != "" || ok {
		log.Warnf("X-Slack-Retry-Num is %s", apiRequest.Headers["X-Slack-Retry-Num"])
		return events.APIGatewayProxyResponse{Body: apiRequest.Body, StatusCode: 200}, nil
	}

	// setup config
	config, err := newConfig()
	if err != nil {
		log.Warn(err.Error())
	}

	// default level is INFO
	if config.Debug {
		log.SetLevel(log.DebugLevel)
	}
	// default log format is ASCII
	if config.LogFormatJSON {
		log.SetFormatter(&log.JSONFormatter{})
	}

	log.Debugf("Header: %+v\n", apiRequest.Headers)
	log.Debugf("Body: %+v\n", apiRequest.Body)

	// decode request
	requestBody, err := model.NewAPIGateWayRequest([]byte(apiRequest.Body), false)
	if err != nil {
		log.Warnf("err: %v\n", err)
		return events.APIGatewayProxyResponse{Body: apiRequest.Body, StatusCode: 200}, nil
	}

	// init model
	program, err := requestBody.ConvertProgram()
	if err != nil {
		log.Warnf("err: %v\n", err)
		return events.APIGatewayProxyResponse{Body: apiRequest.Body, StatusCode: 200}, nil
	}

	// init client
	paizaClient, err := paiza.NewClient()
	if err != nil {
		log.Warn(err.Error())
		return events.APIGatewayProxyResponse{Body: apiRequest.Body, StatusCode: 200}, nil
	}

	// post paiza
	result, err := paizaClient.Request(program)
	if err != nil {
		log.Warnf("err: %v\n", err)
		return events.APIGatewayProxyResponse{Body: apiRequest.Body, StatusCode: 200}, nil
	}

	// TODO:
	slackClient, _ := slack.NewClient("https://hoge/", config.SwiftToken)

	body := slack.SlackRequestBody{}
	body.Token = requestBody.Token
	attachment := slack.Attachment{}
	attachment.Color = "good"
	attachment.Title = "Dummy Title"
	attachment.TitleLink = "https://github.com/TakumiKaribe/multilingo"
	attachment.Text = "```" + result.Stdout + "```"
	body.Attachments = append(body.Attachments, &attachment)
	body.Channel = requestBody.Event.Channel
	body.UserName = requestBody.Event.User

	resp, err := slackClient.Notification(body)
	if err != nil {
		log.Warnf("err: %v\n", err)
		return events.APIGatewayProxyResponse{Body: apiRequest.Body, StatusCode: 200}, nil
	}

	log.Println(resp)

	return events.APIGatewayProxyResponse{Body: apiRequest.Body, StatusCode: 200}, nil
}
