package application

import (
	"context"

	"github.com/TakumiKaribe/multilingo/entity"
	"github.com/TakumiKaribe/multilingo/logger"
	"github.com/aws/aws-lambda-go/events"
)

// LambdaHandler -
func LambdaHandler(ctx context.Context, apiRequest events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// validate retry
	if val, ok := apiRequest.Headers["X-Slack-Retry-Num"]; val != "" || ok {
		logger.Log.Warnf("X-Slack-Retry-Num is %s", apiRequest.Headers["X-Slack-Retry-Num"])
		return events.APIGatewayProxyResponse{Body: apiRequest.Body, StatusCode: 200}, nil
	}

	logger.Log.Infof("Header: %+v\n", apiRequest.Headers)
	logger.Log.Infof("Body: %+v\n", apiRequest.Body)

	// decode request
	requestBody, err := entity.NewAPIGateWayRequestBody([]byte(apiRequest.Body))
	if err != nil {
		logger.Log.Warnf("err: %v\n", err)
		return events.APIGatewayProxyResponse{Body: apiRequest.Body, StatusCode: 500}, nil
	}

	return run(requestBody)
}
