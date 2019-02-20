package slack

import (
	"io"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/pkg/errors"
)

// Client -
type Client struct {
	URL                *url.URL
	HTTPClient         *http.Client
	botUserAccessToken string
	appAuthToken       string
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
func NewClient(urlStr string, botUserAccessToken string, appAuthToken string) (*Client, error) {
	parsedURL, err := url.ParseRequestURI(urlStr)

	if len(botUserAccessToken) == 0 {
		return nil, errors.New("missing  botUserAccessToken")
	}

	if len(appAuthToken) == 0 {
		return nil, errors.New("missing user appAuthToken")
	}

	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse url: %s", urlStr)
	}

	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	return &Client{parsedURL, client, botUserAccessToken, appAuthToken}, nil
}

// NewRequest -
func (c *Client) NewRequest(method string, spath string, body io.Reader) (*http.Request, error) {

	u := *c.URL
	u.Path = path.Join(c.URL.Path, spath)

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "Bearer "+c.botUserAccessToken)

	return req, nil
}

// Notification -
func (c *Client) Notification(body SlackRequestBody) {
	// TODO
	// bodyにappAuthTokenを含める処理を描くお
}
