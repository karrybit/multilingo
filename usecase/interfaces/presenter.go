package interfaces

import "github.com/TakumiKaribe/multilingo/entity/slack"

type Presenter interface {
	ShowResult(*[]*slack.Attachment)
	Challenge()
	LeaveChannel()
	ShowError(error)
}
