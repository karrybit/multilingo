package main

import (
	"github.com/TakumiKaribe/multilingo/model"
	"github.com/TakumiKaribe/multilingo/request/paiza"
	"github.com/TakumiKaribe/multilingo/request/slack"
	log "github.com/sirupsen/logrus"
)

func execDebug() {
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

	body := slack.SlackRequestBody{}
	body.Token = requestBody.Token
	attachment := slack.Attachment{}
	attachment.Color = "good"
	attachment.Title = "Dummy Title"
	attachment.TitleLink = "https://github.com/TakumiKaribe/multilingo"
	attachment.Text = "```" + result.Stdout + "```"
	body.Attachments = append(body.Attachments, &attachment)
	body.Channel = requestBody.Event.Channel
	body.UserName = requestBody.Event.User

	resp, err := slackClient.Notification(body)
	if err != nil {
		log.Warnf("err: %v\n", err)
		return
	}

	log.Println(resp)
}
