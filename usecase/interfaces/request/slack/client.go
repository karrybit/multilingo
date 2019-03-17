package slack

import "github.com/TakumiKaribe/multilingo/entity/slack"

type Client interface {
	Notify(*slack.RequestBody) error
}
