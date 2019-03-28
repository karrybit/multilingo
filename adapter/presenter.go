package adapter

import (
	"github.com/TakumiKaribe/multilingo/entity"
	"github.com/TakumiKaribe/multilingo/entity/config"
	entitySlack "github.com/TakumiKaribe/multilingo/entity/slack"
	requestSlack "github.com/TakumiKaribe/multilingo/infrastructure/request/slack"
	"github.com/TakumiKaribe/multilingo/usecase/interfaces"
	interfaceSlack "github.com/TakumiKaribe/multilingo/usecase/interfaces/request/slack"
)

type Presenter struct {
	client interfaceSlack.Client
	bot    *entitySlack.Bot
	body   *entity.APIGateWayRequestBody
}

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

func (p *Presenter) ShowResult(attachments *[]*entitySlack.Attachment) {
	requestBody := p.makeSlackRequestBody()
	requestBody.Attachments = *attachments

	p.client.Notify(requestBody)
}

func (p *Presenter) ShowError(err error) {
	requestBody := p.makeSlackRequestBody()
	attachments := append([]*entitySlack.Attachment{}, &entitySlack.Attachment{Color: "danger", Title: "[ERROR]", Text: err.Error()})
	requestBody.Attachments = attachments

	p.client.Notify(requestBody)
}

func (p *Presenter) makeSlackRequestBody() *entitySlack.RequestBody {
	return &entitySlack.RequestBody{Token: p.body.Token, Channel: p.body.Event.Channel, UserName: p.bot.Name}
}
