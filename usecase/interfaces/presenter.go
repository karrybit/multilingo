package interfaces

import "github.com/TakumiKaribe/multilingo/entity/slack"

// Presenter -
type Presenter interface {
	ShowResult(*[]*slack.Attachment)
	ShowError(error)
}
