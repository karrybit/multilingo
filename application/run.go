package application

import (
	"github.com/TakumiKaribe/multilingo/adapter"
	"github.com/TakumiKaribe/multilingo/entity"
	"github.com/TakumiKaribe/multilingo/entity/config"
	"github.com/aws/aws-lambda-go/events"
	log "github.com/sirupsen/logrus"
)

func run(requestBody *entity.APIGateWayRequestBody) (events.APIGatewayProxyResponse, error) {
	if requestBody.Challenge != "" {
		response := events.APIGatewayProxyResponse{StatusCode: 200}
		response.Body = requestBody.Challenge
		return response, nil
	}

	// setup config
	config, err := config.NewConfig()
	if err != nil {
		log.Warn(err.Error())
	}

	controller, err := adapter.NewController(config, requestBody)
	if err != nil {
		log.Warn(err.Error())
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
	}

	err = controller.ExecProgram()
	if err != nil {
		log.Warn(err.Error())
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}
