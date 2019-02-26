package slack

import (
	"io"
	"net/http"
	"net/url"
	"path"
	"time"
	"encoding/json"

	"github.com/pkg/errors"
)

// Client -
type Client struct {
	URL                *url.URL
	HTTPClient         *http.Client
	botUserAccessToken string
}

// SlackRequestBody -
type SlackRequestBody struct {
	Token       string        `json:"token"`
	Channel     string        `json:"channel"`
	Attachments []*Attachment `json:"attachments"`
	UserName    string        `json:"username"`
}

// Attachment -
// https://api.slack.com/docs/message-attachments
type Attachment struct {
	Color     string `json:"color"` // good or warning or danger or colorcode
	Title     string `json:"title"`
	TitleLink string `json:"title_link"`
	Text      string `json:"text"`
}

// NewClient Constructor -
func NewClient(botUserAccessToken string) (*Client, error) {
	
	if len(botUserAccessToken) == 0 {
		return nil, errors.New("missing  botUserAccessToken")
	}
	
	client = Client{botUserAccessToken: botUserAccessToken}
	client.URL, _ = url.Parse("https://ntj.slack.com/api/chat.postMessage")
	client.HTTPClient = &http.Client{Timeout: time.Duration(10) * time.Second}
	return &client, nil
}

// newRequest -
func (c *Client) newRequest(method string, body io.Reader) (*http.Request, error) {

	u := *c.URL

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "Bearer "+c.botUserAccessToken)

	return req, nil
}

// Notification -
func (c *Client) Notification(body SlackRequestBody) {
	bodyByte, _ = json.Marshal(body)
	bodyReader := bytes.NewReader(bodyByte)
	
	req, _ := c.newRequest("POST", nil, bodyReader)
    if err != nil {
		// @TODO: リクエストの作成に失敗したとき
        return
    }

    res, _ := c.HTTPClient.Do(req)
}