package running

import (
	"github.com/TakumiKaribe/multilingo/model"
	"github.com/TakumiKaribe/multilingo/request/slack"
	log "github.com/sirupsen/logrus"
)

func noticeError(client *slack.Client, body *model.SlackRequestBody, err error) {
	log.Warnf("err: %v\n", err)
	attachments := append([]*model.Attachment{}, &model.Attachment{Color: "danger", Title: "[ERROR]", Text: err.Error()})
	resp, err := client.Notification(body, &attachments)
	log.Println(resp)
	if err != nil {
		log.Warnf("err: %v\n", err)
		return
	}
}
