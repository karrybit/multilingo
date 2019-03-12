package paiza

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/TakumiKaribe/multilingo/model"
	log "github.com/sirupsen/logrus"
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

// Request -
func (c *Client) Request(program *model.Program) (*model.ExecutionResult, error) {
	status, err := c.execProgram(program)
	if err != nil {
		return nil, err
	}

	// wait execute program until completed
	for isCompleted := false; isCompleted == false; time.Sleep(1 * time.Second) {
		isCompleted, err = c.getStatus(status)
		if err != nil {
			return nil, err
		}
	}

	return c.getResult(status)
}

// ExecProgramRequest is request to execute program
func (c *Client) execProgram(program *model.Program) (*model.Status, error) {
	query := map[string]string{"language": program.Lang, "api_key": "guest", "source_code": program.Program}
	values := url.Values{}
	for k, v := range query {
		values.Add(k, v)
	}

	urlString := c.BaseURL.String() + "create?" + values.Encode()
	req, err := http.NewRequest(http.MethodPost, urlString, nil)
	log.Printf("⚡️  %s\n", urlString)

	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var status model.Status
	err = decoder.Decode(&status)
	if err != nil {
		return nil, err
	}

	return &status, nil
}

// GetStatusRequest is request to get execution status
func (c *Client) getStatus(status *model.Status) (bool, error) {
	query := map[string]string{"id": status.ID, "api_key": "guest"}
	values := url.Values{}
	for k, v := range query {
		values.Add(k, v)
	}

	urlString := c.BaseURL.String() + "get_status?" + values.Encode()

	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	if err != nil {
		return false, err
	}

	log.Printf("⚡️  %s\n", urlString)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(status)
	if err != nil {
		return false, err
	}

	return status.Status == "completed", nil
}

// GetResultRequest is request to get execution result
func (c *Client) getResult(status *model.Status) (*model.ExecutionResult, error) {
	query := map[string]string{"id": status.ID, "api_key": "guest"}
	values := url.Values{}
	for k, v := range query {
		values.Add(k, v)
	}

	urlString := c.BaseURL.String() + "get_details?" + values.Encode()

	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	if err != nil {
		return nil, err
	}

	log.Printf("⚡️  %s\n", urlString)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var result model.ExecutionResult
	err = decoder.Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
