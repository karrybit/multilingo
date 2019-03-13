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
	requestBody, err := model.NewAPIGateWayRequest([]byte{})
	if err != nil {
		log.Warnf("err: %v\n", err)
		return
	}

	// init slack client
	slackClient, err := slack.NewClient("https://hoge/", "")
	if err != nil {
		log.Warnf("err: %v\n", err)
		return
	}

	// init model
	program, err := requestBody.ConvertProgram()
	if err != nil {
		noticeError(slackClient, requestBody.ConvertSlackRequestBody(), err)
	}

	// init paiza client
	paizaClient, err := paiza.NewClient()
	if err != nil {
		noticeError(slackClient, requestBody.ConvertSlackRequestBody(), err)
	}

	// post paiza
	result, err := paizaClient.Request(program)
	if err != nil {
		noticeError(slackClient, requestBody.ConvertSlackRequestBody(), err)
	}

	log.Printf("%+v", result)

	resp, err := slackClient.Notification(requestBody.ConvertSlackRequestBody(), result.MakeAttachments())
	if err != nil {
		log.Warnf("err: %v\n", err)
		return
	}

	log.Println(resp)
}
