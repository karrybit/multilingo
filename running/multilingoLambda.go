package running

import (
	"context"

	"github.com/TakumiKaribe/multilingo/config"
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
	config, err := config.NewConfig()
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

	slackClient, err := slack.NewClient("https://hoge/", config.SwiftToken)
	if err != nil {
		noticeError(slackClient, requestBody.ConvertSlackRequestBody(), err)
		return events.APIGatewayProxyResponse{Body: apiRequest.Body, StatusCode: 200}, nil
	}

	// init model
	program, err := requestBody.ConvertProgram()
	if err != nil {
		noticeError(slackClient, requestBody.ConvertSlackRequestBody(), err)
		return events.APIGatewayProxyResponse{Body: apiRequest.Body, StatusCode: 200}, nil
	}

	// init client
	paizaClient, err := paiza.NewClient()
	if err != nil {
		noticeError(slackClient, requestBody.ConvertSlackRequestBody(), err)
		return events.APIGatewayProxyResponse{Body: apiRequest.Body, StatusCode: 200}, nil
	}

	// post paiza
	result, err := paizaClient.Request(program)
	if err != nil {
		noticeError(slackClient, requestBody.ConvertSlackRequestBody(), err)
		return events.APIGatewayProxyResponse{Body: apiRequest.Body, StatusCode: 200}, nil
	}

	resp, err := slackClient.Notification(requestBody.ConvertSlackRequestBody(), result.MakeAttachments())
	if err != nil {
		noticeError(slackClient, requestBody.ConvertSlackRequestBody(), err)
		return events.APIGatewayProxyResponse{Body: apiRequest.Body, StatusCode: 200}, nil
	}

	log.Println(resp)
	return events.APIGatewayProxyResponse{Body: apiRequest.Body, StatusCode: 200}, nil
}
