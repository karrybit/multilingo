package interfaces

// UseCase -
type UseCase interface {
	ExecProgram(string, string) error
	Challenge()
	Kick()
}
