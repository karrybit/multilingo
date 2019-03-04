package paiza

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/TakumiKaribe/multilingo/request/paiza/model"
)

// Client -
type Client struct {
	URL        *url.URL
	HTTPClient *http.Client
}

// NewClient Constructor -
func NewClient() (*Client, error) {
	client := Client{HTTPClient: &http.Client{Timeout: time.Duration(10) * time.Second}}
	client.URL, _ = url.Parse("http://api.paiza.io:80/runners/")

	return &client, nil
}

// StatusResult -
type StatusResult struct {
	Response model.Status
	Err      error
}

// ExecProgramRequest is request to execute program
func (c *Client) ExecProgram(query map[string]string, ch chan<- StatusResult) {
	result := StatusResult{}

	bodyByte, _ := json.Marshal(query)
	bodyReader := bytes.NewReader(bodyByte)

	req, err := http.NewRequest("POST", c.URL.String(), bodyReader)
	// TODO: use loglus
	log.Printf("⚡️  %s\n", c.URL.String())
	if err != nil {
		result.Err = err
		ch <- result
	}

	defer req.Body.Close()

	resp, err := c.HTTPClient.Do(req)

	decoder := json.NewDecoder(resp.Body)
	var status model.Status
	err = decoder.Decode(&status)
	if err != nil {
		result.Err = err
		ch <- result
	}

	result.Response = status
	ch <- result
}
