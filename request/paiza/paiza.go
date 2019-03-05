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
	BaseURL    *url.URL
	HTTPClient *http.Client
}

// NewClient Constructor -
func NewClient() (*Client, error) {
	client := Client{HTTPClient: &http.Client{Timeout: time.Duration(10) * time.Second}}
	client.BaseURL, _ = url.Parse("http://api.paiza.io:80/runners/")

	return &client, nil
}

// StatusResult -
type StatusResult struct {
	Response model.Status
	Err      error
}

// ExecProgramRequest is request to execute program
func (c *Client) ExecProgram(lang string, program string, ch chan<- StatusResult) {
	result := StatusResult{}

	query := map[string]string{"language": lang, "api_key": "guest", "source_code": program}
	bodyByte, _ := json.Marshal(query)
	bodyReader := bytes.NewReader(bodyByte)

	urlString := c.BaseURL.String() + "create"

	req, err := http.NewRequest("POST", urlString, bodyReader)
	// TODO: use loglus
	log.Printf("⚡️  %s\n", urlString)

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

// GetStatusRequest is request to get execution status
func (c *Client) GetStatusRequest(id string, ch chan<- StatusResult) {
	result := StatusResult{}

	query := map[string]string{"id": id, "api_key": "guest"}
	bodyByte, _ := json.Marshal(query)
	bodyReader := bytes.NewReader(bodyByte)

	urlString := c.BaseURL.String() + "get_status"

	req, err := http.NewRequest("GET", urlString, bodyReader)
	// TODO: use loglus
	log.Printf("⚡️  %s\n", urlString)

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

// ExecutionResult -
type ExecutionResult struct {
	Response model.ExecutionResult
	Err      error
}

// GetResultRequest is request to get execution result
func (c *Client) GetResultRequest(id string, ch chan<- ExecutionResult) {
	result := ExecutionResult{}

	query := map[string]string{"id": id, "api_key": "guest"}
	bodyByte, _ := json.Marshal(query)
	bodyReader := bytes.NewReader(bodyByte)

	urlString := c.BaseURL.String() + "get_details"

	req, err := http.NewRequest("GET", urlString, bodyReader)
	// TODO: use loglus
	log.Printf("⚡️  %s\n", urlString)

	if err != nil {
		result.Err = err
		ch <- result
	}

	defer req.Body.Close()

	resp, err := c.HTTPClient.Do(req)

	decoder := json.NewDecoder(resp.Body)
	var executionResult model.ExecutionResult
	err = decoder.Decode(&executionResult)
	if err != nil {
		result.Err = err
		ch <- result
	}

	result.Response = executionResult
	ch <- result
}
