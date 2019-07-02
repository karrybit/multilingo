package adapter

import (
	"multilingo/entity"
	"multilingo/entity/config"
	"multilingo/logger"
	"multilingo/usecase/interactor"
	"multilingo/usecase/interfaces"
)

// Controller is controller by clean architecture
type Controller struct {
	body    *entity.APIGateWayRequestBody
	useCase interfaces.UseCase
}

// NewController is initialize Controller
func NewController(requestBody *entity.APIGateWayRequestBody) (*Controller, error) {
	presenter, err := NewPresenter(requestBody)
	if err != nil {
		logger.Log.Warn(err.Error())
		return nil, err
	}
	return &Controller{body: requestBody, useCase: interactor.NewInteractor(presenter)}, nil
}

// ExecProgram is to exec program by post to paiza
func (c *Controller) ExecProgram() error {
	// init bot_info
	bot, err := config.SharedConfig.NewBotInfo(c.body.APIAppID)
	if err != nil {
		logger.Log.Warn(err.Error())
		return err
	}

	return c.useCase.ExecProgram(bot.Language, c.body.Event.Text)
}
