package adapter

import (
	"multilingo/entity"
	"multilingo/entity/config"
	entitySlack "multilingo/entity/slack"
	requestSlack "multilingo/infrastructure/request/slack"
	"multilingo/usecase/interfaces"
	interfaceSlack "multilingo/usecase/interfaces/request/slack"
)

// Presenter -
type Presenter struct {
	client interfaceSlack.Client
	bot    *entitySlack.Bot
	body   *entity.APIGateWayRequestBody
}

// NewPresenter -
func NewPresenter(body *entity.APIGateWayRequestBody) (interfaces.Presenter, error) {
	presenter := Presenter{}
	bot, err := config.SharedConfig.NewBotInfo(body.APIAppID)
	if err != nil {
		return nil, err
	}
	presenter.bot = bot
	presenter.body = body
	client, err := requestSlack.NewClient(config.SharedConfig.SlackPath, bot.Token)
	if err != nil {
		return nil, err
	}
	presenter.client = client
	return &presenter, nil
}

// ShowResult -
func (p *Presenter) ShowResult(attachments *[]*entitySlack.Attachment) {
	requestBody := p.makeSlackRequestBody()
	requestBody.Attachments = *attachments

	p.client.Notify(requestBody)
}

// ShowError -
func (p *Presenter) ShowError(err error) {
	requestBody := p.makeSlackRequestBody()
	attachments := append([]*entitySlack.Attachment{}, &entitySlack.Attachment{Color: "danger", Title: "[ERROR]", Text: err.Error()})
	requestBody.Attachments = attachments

	p.client.Notify(requestBody)
}

func (p *Presenter) makeSlackRequestBody() *entitySlack.RequestBody {
	return &entitySlack.RequestBody{Token: p.body.Token, Channel: p.body.Event.Channel, UserName: p.bot.Name}
}
