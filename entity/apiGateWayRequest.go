package entity

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// APIGateWayRequestBody -
type APIGateWayRequestBody struct {
	Token     string `json:"token"`
	TeamID    string `json:"team_id"`
	APIAppID  string `json:"api_app_id"`
	Event     Event  `json:"event"`
	Challenge string `json:"challenge"`
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

// NewAPIGateWayRequestBody -
func NewAPIGateWayRequestBody(bytes []byte) (*APIGateWayRequestBody, error) {
	var request APIGateWayRequestBody
	err := json.Unmarshal(bytes, &request)
	if err != nil {
		return nil, errors.Wrap(err, "failed unmarshal APIGateWayRequestBody")
	}
	return &request, nil
}
