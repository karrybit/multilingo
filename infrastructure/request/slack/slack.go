package slack

import (
	"bytes"
	"encoding/json"
	"net/url"

	"github.com/TakumiKaribe/multilingo/entity"
	"github.com/TakumiKaribe/multilingo/usecase/interfaces"

	"github.com/pkg/errors"
)

const (
	postMessage = "/api/chat.postMessage"
)

// Client -
type Client struct {
	url                *url.URL
	requester          *interfaces.Reqeuster
	botUserAccessToken string
}

// NewClient Constructor -
func NewClient(host string, botUserAccessToken string) (*Client, error) {
	if len(botUserAccessToken) == 0 {
		// TODO: use multilingo error
		return nil, errors.New("missing  botUserAccessToken")
	}

	client := Client{botUserAccessToken: botUserAccessToken, requester: interfaces.NewRequester()}
	client.url, _ = url.Parse(host + postMessage)

	return &client, nil
}

// Notification -
func (c *Client) Notify(requestBody *entity.SlackRequestBody) error {
	bodyByte, _ := json.Marshal(requestBody)
	bodyReader := bytes.NewReader(bodyByte)

	header := map[string]string{}
	header["Authorization"] = "Bearer " + c.botUserAccessToken
	header["Content-Type"] = "application/json; charset=UTF-8"
	body, err := c.requester.Request(interfaces.Post, c.url.String(), bodyReader, header)
	if err != nil {
		return err
	}

	defer body.Close()
	return nil
}
