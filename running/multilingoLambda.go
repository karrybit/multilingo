package running

import (
	"context"

	"github.com/TakumiKaribe/multilingo/model"
	"github.com/aws/aws-lambda-go/events"
	log "github.com/sirupsen/logrus"
)

// LambdaHandler -
func LambdaHandler(ctx context.Context, apiRequest events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.SetFormatter(&log.JSONFormatter{})

	// validate retry
	if val, ok := apiRequest.Headers["X-Slack-Retry-Num"]; val != "" || ok {
		log.Warnf("X-Slack-Retry-Num is %s", apiRequest.Headers["X-Slack-Retry-Num"])
		return events.APIGatewayProxyResponse{Body: apiRequest.Body, StatusCode: 200}, nil
	}

	log.Debugf("Header: %+v\n", apiRequest.Headers)
	log.Debugf("Body: %+v\n", apiRequest.Body)

	// decode request
	requestBody, err := model.NewAPIGateWayRequestBody([]byte(apiRequest.Body))
	if err != nil {
		log.Warnf("err: %v\n", err)
		return events.APIGatewayProxyResponse{Body: apiRequest.Body, StatusCode: 200}, nil
	}

	return run(requestBody)
}
