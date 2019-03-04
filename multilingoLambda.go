package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/TakumiKaribe/multilingo/parserawtext"
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

func HelloLambdaHandler(ctx context.Context, apiRequest events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
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

	// TODO:
	/*
		// default level is INFO
		if conf.Debug {
			log.SetLevel(log.DebugLevel)
		}
		// default log format is ASCII
		if conf.LogFormatJson {
			log.SetFormatter(&log.JSONFormatter{})
		}
	*/

	// parse program
	text, err := parserawtext.Parse(requestBody.Event.Text)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return events.APIGatewayProxyResponse{Body: apiRequest.Body, StatusCode: 400}, nil
	}

	// post paiza
	status, err := execProgram(lang, text)
	if err != err {
		log.Warn(err.Error())
	}

	result, err := getResult(status)
	if err != err {
		log.Warn(err.Error())
	}

	// TODO:
	_ = result

	// TODO:
	client, _ := slack.NewClient("host", "token")

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

	client.Notification(body)

	return events.APIGatewayProxyResponse{Body: apiRequest.Body, StatusCode: 222}, nil
}
