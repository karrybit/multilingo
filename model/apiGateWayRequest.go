package model

import (
	"encoding/json"

	"github.com/TakumiKaribe/multilingo/parsetext"
)

// APIGateWayRequest -
type APIGateWayRequest struct {
	Token    string `json:"token"`
	TeamID   string `json:"team_id"`
	APIAppID string `json:"api_app_id"`
	Event    Event  `json:"event"`
}

// Event -
type Event struct {
	ClientMsgID    string `json:"client_msg_id"`
	EventType      string `json:"type"`
	Text           string `json:"text"`
	User           string `json:"user"`
	Timestamp      string `json:"ts"`
	Channel        string `json:"channel"`
	EventTimestamp string `json:"event_ts"`
}

// NewAPIGateWayRequest -
func NewAPIGateWayRequest(bytes []byte, isDebug bool) (*APIGateWayRequest, error) {
	if isDebug {
		event := Event{Text: "<@debug>```print(114514)```", User: "Swift", Channel: "#bot"}
		request := APIGateWayRequest{Token: "Token", APIAppID: "apiAppID", Event: event}
		return &request, nil
	}

	var request APIGateWayRequest
	err := json.Unmarshal(bytes, &request)
	if err != nil {
		return nil, err
	}
	return &request, nil
}

// ConvertProgram -
func (r *APIGateWayRequest) ConvertProgram() (*Program, error) {
	lang, err := parsetext.LookUpLanguage(r.APIAppID)
	if err != nil {
		return nil, err
	}

	program, err := parsetext.Parse(r.Event.Text)
	if err != nil {
		return nil, err
	}

	return &Program{Lang: lang, Program: program}, nil
}
