package interfaces

import "multilingo/entity/slack"

// Presenter -
type Presenter interface {
	ShowResult(*[]*slack.Attachment)
	ShowError(error)
}
