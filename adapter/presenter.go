package adapter

import (
	"github.com/TakumiKaribe/multilingo/entity"
	"github.com/TakumiKaribe/multilingo/entity/config"
	"github.com/TakumiKaribe/multilingo/infrastructure/request/slack"
	"github.com/TakumiKaribe/multilingo/usecase/interfaces"
)

type Presenter struct {
	client *slack.Client
	bot    *entity.Bot
	body   *entity.APIGateWayRequestBody
}

func NewPresenter(config *config.Config, body *entity.APIGateWayRequestBody) (interfaces.Presenter, error) {
	presenter := Presenter{}
	bot, err := config.NewBotInfo(body.APIAppID)
	if err != nil {
		return nil, err
	}
	presenter.bot = bot
	presenter.body = body
	client, err := slack.NewClient(config.SlackPath, bot.Token)
	if err != nil {
		return nil, err
	}
	presenter.client = client
	return &presenter, nil
}

func (p *Presenter) ShowResult(result *entity.ExecutionResult) {
	requestBody := p.body.ConvertSlackRequestBody(p.bot)
	requestBody.Attachments = *result.MakeAttachments()

	p.client.Notify(requestBody)
}

func (p *Presenter) Challenge() {
}

func (p *Presenter) LeaveChannel() {
}

func (p *Presenter) ShowError(err error) {
	requestBody := p.body.ConvertSlackRequestBody(p.bot)
	attachments := append([]*entity.Attachment{}, &entity.Attachment{Color: "danger", Title: "[ERROR]", Text: err.Error()})
	requestBody.Attachments = attachments

	p.client.Notify(requestBody)
}
