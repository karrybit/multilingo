package running

import (
	"github.com/TakumiKaribe/multilingo/config"
	"github.com/TakumiKaribe/multilingo/model"
	"github.com/TakumiKaribe/multilingo/parsetext"
	"github.com/TakumiKaribe/multilingo/request/paiza"
	"github.com/TakumiKaribe/multilingo/request/slack"
	"github.com/aws/aws-lambda-go/events"
	log "github.com/sirupsen/logrus"
)

func run(requestBody *model.APIGateWayRequestBody) (events.APIGatewayProxyResponse, error) {
	// setup config
	config, err := config.NewConfig()
	if err != nil {
		log.Warn(err.Error())
	}

	// init bot_info
	bot, err := config.NewBotInfo(requestBody.APIAppID)
	if err != nil {
		log.Warn(err.Error())
	}

	// init slack client
	slackClient, err := slack.NewClient(config.SlackPath, bot.Token)
	if err != nil {
		noticeError(slackClient, requestBody.ConvertSlackRequestBody(bot), err)
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
	}

	// init paiza client
	paizaClient := paiza.NewClient()

	// parse program
	program, err := parsetext.Parse(requestBody.Event.Text)
	if err != nil {
		noticeError(slackClient, requestBody.ConvertSlackRequestBody(bot), err)
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
	}

	// post paiza
	result, err := paizaClient.Request(&model.Program{Lang: bot.Language, Program: program})
	if err != nil {
		noticeError(slackClient, requestBody.ConvertSlackRequestBody(bot), err)
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
	}

	// post slack
	body := requestBody.ConvertSlackRequestBody(bot)
	body.Attachments = *result.MakeAttachments()
	resp, err := slackClient.Notification(body)
	if err != nil {
		noticeError(slackClient, requestBody.ConvertSlackRequestBody(bot), err)
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
	}

	log.Println(resp)
	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}
