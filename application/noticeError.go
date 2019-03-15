package application

import (
	"github.com/TakumiKaribe/multilingo/entity"
	"github.com/TakumiKaribe/multilingo/infrastructure/request/slack"
	log "github.com/sirupsen/logrus"
)

func noticeError(client *slack.Client, body *entity.SlackRequestBody, err error) {
	log.Warnf("err: %v\n", err)
	attachments := append([]*entity.Attachment{}, &entity.Attachment{Color: "danger", Title: "[ERROR]", Text: err.Error()})
	body.Attachments = attachments
	resp, err := client.Notification(body)
	log.Println(resp)
	if err != nil {
		log.Warnf("err: %v\n", err)
		return
	}
}
