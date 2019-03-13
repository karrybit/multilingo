package model

import (
	"encoding/json"
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
func NewAPIGateWayRequest(bytes []byte) (*APIGateWayRequest, error) {
	var request APIGateWayRequest
	err := json.Unmarshal(bytes, &request)
	if err != nil {
		return nil, err
	}
	return &request, nil
}

// ConvertSlackRequestBody -
func (r *APIGateWayRequest) ConvertSlackRequestBody() *SlackRequestBody {
	return &SlackRequestBody{Token: r.Token, Channel: r.Event.Channel, UserName: r.Event.User}
}
