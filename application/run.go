package application

import (
	"multilingo/adapter"
	"multilingo/entity"
	"multilingo/logger"
	"github.com/aws/aws-lambda-go/events"
)

func run(requestBody *entity.APIGateWayRequestBody) (events.APIGatewayProxyResponse, error) {
	if requestBody.Challenge != "" {
		response := events.APIGatewayProxyResponse{StatusCode: 200}
		response.Body = requestBody.Challenge
		return response, nil
	}

	controller, err := adapter.NewController(requestBody)
	if err != nil {
		logger.Log.Warn(err.Error())
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
	}

	err = controller.ExecProgram()
	if err != nil {
		logger.Log.Warn(err.Error())
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}
