package multilingoerror

import (
	"fmt"

	"github.com/pkg/errors"
)

// ErrorType -
type ErrorType int

const (
	// Config
	NotFoundConfig = iota
	NewBotCorrespondingToID
	MissingNotUserAccessToken

	// UseCase
	ParseProgram

	// Request
	FailedRequest
	FailedPaizaCreateRequest
	FailedPaizaStatusRequest

	// Decode
	DecodeAPIGateWayRequest
	DecodePaizaStatus
	DecodePaizaResult
)

// New -
func New(et ErrorType, actual string, expected string) error {
	switch et {
	case NotFoundConfig:
		return fmt.Errorf("No config corresponding to %s was found", actual)
	case NewBotCorrespondingToID:
		return fmt.Errorf("No bot corresponding to %s was found", actual)
	case MissingNotUserAccessToken:
		return errors.New("missing  botUserAccessToken")

	case ParseProgram:
		return errors.New("failed to parse post messages")

	case DecodePaizaStatus:
		return errors.New("failed to decode paiza.Status")
	case DecodePaizaResult:
		return errors.New("failed to decode paiza.Result")

	default:
		return errors.New("unknown error")
	}
}

// Wrap -
func Wrap(et ErrorType, err error) error {
	switch et {
	case FailedRequest:
		return errors.Wrap(err, "failed httpClient.Do")
	case FailedPaizaCreateRequest:
		return errors.Wrap(err, "failed paiza create request")
	case FailedPaizaStatusRequest:
		return errors.Wrap(err, "failed paiza status request")

	case DecodeAPIGateWayRequest:
		return errors.Wrap(err, "failed unmarshal APIGateWayRequestBody")

	default:
		return errors.New("unknown error")
	}
}
