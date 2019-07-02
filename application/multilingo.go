package application

import (
	"context"

	"multilingo/entity"
	"multilingo/logger"
	"github.com/aws/aws-lambda-go/events"
)

// LambdaHandler -
func LambdaHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// validate retry
	if val, ok := request.Headers["X-Slack-Retry-Num"]; val != "" || ok {
		logger.Log.Warnf("X-Slack-Retry-Num is %s", request.Headers["X-Slack-Retry-Num"])
		return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
	}

	logger.Log.Infof("Request: %+v\n", request)

	// decode request
	requestBody, err := entity.NewAPIGateWayRequestBody([]byte(request.Body))
	if err != nil {
		logger.Log.Warnf("err: %v\n", err)
		return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 500}, nil
	}

	return run(requestBody)
}
