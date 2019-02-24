package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/TakumiKaribe/MultilinGo/logger"
	"github.com/TakumiKaribe/MultilinGo/model"
	"github.com/TakumiKaribe/MultilinGo/parserawtext"
	"github.com/TakumiKaribe/MultilinGo/request"
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

	// parse program
	text, err := parserawtext.Parse(requestBody.Event.Text)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return events.APIGatewayProxyResponse{Body: apiRequest.Body, StatusCode: 400}, nil
	}

	// post paiza
	status := execProgram(lang, text)
	getResult(status)

	return events.APIGatewayProxyResponse{Body: apiRequest.Body, StatusCode: 200}, nil
}

func main() {
	if false {
		lambda.Start(HelloLambdaHandler)
	}
}

func execProgram(lang string, program string) model.Status {
	query := map[string]string{"language": lang, "api_key": "guest", "source_code": program}

	ch := make(chan request.StatusResult)
	go request.ExecProgramRequest(query, ch)

	result := <-ch

	if result.Err != nil {
		fmt.Println(result.Err)
		return model.Status{}
	}

	logger.PrintFields(&result.Response)

	return result.Response
}

func getResult(status model.Status) {
	query := map[string]string{"id": status.ID, "api_key": "guest"}

	// wait execute program until completed
	for status := "runnig"; status != "completed"; time.Sleep(1 * time.Second) {
		ch := make(chan request.StatusResult)
		go request.GetStatusRequest(query, ch)
		statusResult := <-ch
		status = statusResult.Response.Status
	}

	ch := make(chan request.ExecutionResult)
	go request.GetResultRequest(query, ch)

	detailResult := <-ch

	if detailResult.Err != nil {
		fmt.Println(detailResult.Err)
		return
	}

	logger.PrintFields(&detailResult.Response)
}
