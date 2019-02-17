package model

import (
	"log"
	"net/http"
	"net/url"
	"strings"
)

// SlackRequestData -
type SlackRequestData struct {
	Token    string `json:"token"`
	Channel  string `json:"channel"`
	Text     string `json:"text"`
	UserName string `json:"username"`
}

// Notification Method -
func (data *SlackRequestData) Notification() {
	slackURL := "https://ntj.slack.com/api/chat.postMessage"
	values := url.Values{}
	values.Add("token", data.Token)
	values.Add("channel", data.Channel)
	values.Add("text", data.Text)
	values.Add("username", data.UserName)

	resp, err := http.NewRequest(
		"POST",
		slackURL,
		strings.NewReader(values.Encode()),
	)

	log.Printf("⚡️  %s\n", resp)
	log.Printf("⚡️  %s\n", err)
}

// Log is standard output of all properties
func (dr *SlackRequestData) Log() {}
