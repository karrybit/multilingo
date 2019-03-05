package main

import (
	"time"

	"github.com/TakumiKaribe/multilingo/parsetext"
	"github.com/TakumiKaribe/multilingo/request/paiza"
	"github.com/TakumiKaribe/multilingo/request/slack"
	log "github.com/sirupsen/logrus"
)

func execDebug(appID string, program string, token string, channel string, user string) {
	// setup config
	config, err := newConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	// default level is INFO
	if config.Debug {
		log.SetLevel(log.DebugLevel)
	}
	// default log format is ASCII
	if config.LogFormatJSON {
		log.SetFormatter(&log.JSONFormatter{})
	}

	// look up language type
	lang, err := debugLookUpLanguage(appID)
	if err != nil {
		log.Warnf("err: %v\n", err)
	}

	// parse program
	text, err := parsetext.Parse(program)
	if err != nil {
		log.Warnf("err: %v\n", err)
	}

	// post paiza
	paizaClient, err := paiza.NewClient()
	if err != nil {
		log.Warn(err.Error())
	}

	execCh := make(chan paiza.StatusResult)
	go paizaClient.ExecProgram(lang, text, execCh)
	execStatus := <-execCh
	if execStatus.Err != nil {
		log.Warn(execStatus.Err.Error())
	}
	log.Debug(&execStatus.Response)

	// wait execute program until completed
	for status := "runnig"; status != "completed"; time.Sleep(1 * time.Second) {
		statusCh := make(chan paiza.StatusResult)
		go paizaClient.GetStatusRequest(execStatus.Response.ID, statusCh)
		statusResult := <-statusCh
		status = statusResult.Response.Status
	}

	detailCh := make(chan paiza.ExecutionResult)
	go paizaClient.GetResultRequest(execStatus.Response.ID, detailCh)
	executionResult := <-detailCh

	if executionResult.Err != nil {
		log.Warn(executionResult.Err.Error())
	}

	// TODO:
	slackClient, _ := slack.NewClient("host", "token")

	body := slack.SlackRequestBody{}
	body.Token = token
	attachment := slack.Attachment{}
	attachment.Color = "good"
	attachment.Title = "Dummy Title"
	attachment.TitleLink = "https://github.com/TakumiKaribe/multilingo"
	attachment.Text = "```" + executionResult.Response.Stdout + "```"
	body.Attachments = append(body.Attachments, &attachment)
	body.Channel = channel
	body.UserName = user

	slackClient.Notification(body)
}
