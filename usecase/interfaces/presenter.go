package interfaces

import "github.com/TakumiKaribe/multilingo/entity/paiza"

type Presenter interface {
	ShowResult(*paiza.Result)
	Challenge()
	LeaveChannel()
	ShowError(error)
}
