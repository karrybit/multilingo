package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/TakumiKaribe/multilingo/parserawtext"
	"github.com/TakumiKaribe/multilingo/request/paiza"
	"github.com/TakumiKaribe/multilingo/request/slack"
	"github.com/aws/aws-lambda-go/events"
	log "github.com/sirupsen/logrus"
)

// TODO: naming
type APIGateWayRequest struct {
	Token    string `json:"token"`
	TeamID   string `json:"team_id"`
	ApiAppID string `json:"api_app_id"`
	Event    Event  `json:"event"`
}

type Event struct {
	ClientMsgId    string `json:"client_msg_id"`
	EventType      string `json:"type"`
	Text           string `json:"text"`
	User           string `json:"user"`
	Timestamp      string `json:"ts"`
	Channel        string `json:"channel"`
	EventTimestamp string `json:"event_ts"`
}

func LambdaHandler(ctx context.Context, apiRequest events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Body: %v\n", apiRequest.Headers)
	fmt.Printf("Body: %v\n", apiRequest.Body)

	// parse json
	var requestBody APIGateWayRequest
	err := json.Unmarshal([]byte(apiRequest.Body), &requestBody)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return events.APIGatewayProxyResponse{Body: apiRequest.Body, StatusCode: 400}, nil
	}

	// look up language type
	lang, err := lookUpLanguage(&requestBody)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return events.APIGatewayProxyResponse{Body: apiRequest.Body, StatusCode: 400}, nil
	}

	config, err := newConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	// default level is INFO
	if config.Debug {
		log.SetLevel(log.DebugLevel)
	}
	// default log format is ASCII
	if config.LogFormatJson {
		log.SetFormatter(&log.JSONFormatter{})
	}

	// parse program
	text, err := parserawtext.Parse(requestBody.Event.Text)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return events.APIGatewayProxyResponse{Body: apiRequest.Body, StatusCode: 400}, nil
	}

	// post paiza
	paizaClient, err := paiza.NewClient()
	if err != nil {
		log.Warn(err.Error())
		return events.APIGatewayProxyResponse{Body: apiRequest.Body, StatusCode: 400}, nil
	}

	execCh := make(chan paiza.StatusResult)
	go paizaClient.ExecProgram(lang, text, execCh)
	execStatus := <-execCh
	if execStatus.Err != nil {
		log.Warn(execStatus.Err.Error())
	}
	log.Debug(&execStatus.Response)

	// wait execute program until completed
	for status := "runnig"; status != "completed"; time.Sleep(1 * time.Second) {
		statusCh := make(chan paiza.StatusResult)
		go paizaClient.GetStatusRequest(execStatus.Response.ID, statusCh)
		statusResult := <-statusCh
		status = statusResult.Response.Status
	}

	detailCh := make(chan paiza.ExecutionResult)
	go paizaClient.GetResultRequest(execStatus.Response.ID, detailCh)
	executionResult := <-detailCh

	if executionResult.Err != nil {
		log.Warn(executionResult.Err.Error())
	}

	// TODO:
	slackClient, _ := slack.NewClient("host", "token")

	body := slack.SlackRequestBody{}
	body.Token = requestBody.Token
	attachment := slack.Attachment{}
	attachment.Color = "good"
	attachment.Title = "Dummy Title"
	attachment.TitleLink = "https://github.com/TakumiKaribe/multilingo"
	// TODO:
	// attachment.Text = result.Response.Stdout
	attachment.Text = "ここに実行結果が入るよ"
	body.Attachments = append(body.Attachments, &attachment)
	body.Channel = requestBody.Event.Channel
	body.UserName = requestBody.Event.User

	slackClient.Notification(body)

	return events.APIGatewayProxyResponse{Body: apiRequest.Body, StatusCode: 222}, nil
}
