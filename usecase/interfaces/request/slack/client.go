package slack

import "multilingo/entity/slack"

type Client interface {
	Notify(*slack.RequestBody) error
}
