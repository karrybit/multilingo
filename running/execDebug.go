package running

import (
	"github.com/TakumiKaribe/multilingo/model"
	"github.com/TakumiKaribe/multilingo/request/paiza"
	"github.com/TakumiKaribe/multilingo/request/slack"
	log "github.com/sirupsen/logrus"
)

func ExecDebug() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.JSONFormatter{})

	// decode request
	requestBody, err := model.NewAPIGateWayRequest([]byte{}, true)
	if err != nil {
		log.Warnf("err: %v\n", err)
		return
	}

	// init model
	program, err := requestBody.ConvertProgram()
	if err != nil {
		log.Warnf("err: %v\n", err)
		return
	}

	// init client
	paizaClient, err := paiza.NewClient()
	if err != nil {
		log.Warn(err.Error())
		return
	}

	// post paiza
	result, err := paizaClient.Request(program)
	if err != nil {
		log.Warnf("err: %v\n", err)
		return
	}

	log.Printf("%+v", result)

	// TODO:
	slackClient, err := slack.NewClient("https://hoge/", "BotUserAccessToken")
	if err != nil {
		log.Warnf("err: %v\n", err)
		return
	}

	resp, err := slackClient.Notification(requestBody.ConvertSlackRequestBody(), result)
	if err != nil {
		log.Warnf("err: %v\n", err)
		return
	}

	log.Println(resp)
}
