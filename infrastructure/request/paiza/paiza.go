package paiza

import (
	"encoding/json"
	"net/url"
	"strings"
	"time"

	"github.com/TakumiKaribe/multilingo/entity/paiza"
	"github.com/TakumiKaribe/multilingo/usecase/interfaces"
)

const (
	baseURL   = "http://api.paiza.io:80/runners/"
	create    = "create?"
	getStatus = "get_status?"
	getDetail = "get_details?"
)

type client struct {
	requester     *interfaces.Reqeuster
	baseURL       *url.URL
	defaultParams url.Values
}

func NewClient() interfaces.PaizaClient {
	client := client{requester: interfaces.NewRequester()}
	client.baseURL, _ = url.Parse(baseURL)
	client.defaultParams = url.Values{}
	client.defaultParams.Add("api_key", "guest")
	return &client
}

func (c *client) Request(language string, program string) (*paiza.Result, error) {
	status, err := c.create(language, program)
	if err != nil {
		return nil, err
	}

	for ; status.Status != "completed"; time.Sleep(1 * time.Second) {
		status, err = c.getStatus(status)
		if err != nil {
			return nil, err
		}
	}

	return c.getDetail(status.ID)
}

func (c *client) create(language string, program string) (*paiza.Status, error) {
	params := c.defaultParams
	params.Add("language", language)
	params.Add("source_code", program)

	urlString := strings.Join([]string{c.baseURL.String(), create, params.Encode()}, "")

	body, err := c.requester.Request(interfaces.Post, urlString, nil, map[string]string{})
	if err != nil {
		return nil, err
	}

	defer body.Close()
	decoder := json.NewDecoder(body)
	var status paiza.Status
	err = decoder.Decode(&status)
	if err != nil {
		return nil, err
	}

	return &status, nil
}

func (c *client) getStatus(status *paiza.Status) (*paiza.Status, error) {
	params := c.defaultParams
	params.Add("id", status.ID)

	urlString := strings.Join([]string{c.baseURL.String(), getStatus, params.Encode()}, "")
	body, err := c.requester.Request(interfaces.Get, urlString, nil, map[string]string{})

	defer body.Close()
	decoder := json.NewDecoder(body)
	err = decoder.Decode(status)
	if err != nil {
		return nil, err
	}

	return status, nil
}

func (c *client) getDetail(id string) (*paiza.Result, error) {
	params := c.defaultParams
	params.Add("id", id)

	urlString := strings.Join([]string{c.baseURL.String(), getDetail, params.Encode()}, "")
	body, err := c.requester.Request(interfaces.Get, urlString, nil, map[string]string{})

	defer body.Close()
	decoder := json.NewDecoder(body)

	var result paiza.Result
	err = decoder.Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
