package interfaces

import "github.com/TakumiKaribe/multilingo/entity"

type Presenter interface {
	ShowResult(*entity.ExecutionResult)
	Challenge()
	LeaveChannel()
	ShowError(error)
}
