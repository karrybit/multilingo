package multilingoerror

import (
	"errors"
	"fmt"
)

type ErrorType int

const (
	NotFoundConfig = iota
)

func New(et ErrorType, actual string, expected string) error {
	switch et {
	case NotFoundConfig:
		return fmt.Errorf("No config corresponding to %s was found", actual)
	default:
		return errors.New("unknown error")
	}
}
