package paiza

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/TakumiKaribe/multilingo/entity"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type endPoint string

const (
	create    endPoint = "create?"
	getStatus endPoint = "get_status?"
	getDetail endPoint = "get_details?"
)

// Client -
type Client struct {
	BaseURL    *url.URL
	HTTPClient *http.Client
}

// NewClient Constructor -
func NewClient() *Client {
	client := Client{HTTPClient: &http.Client{Timeout: time.Duration(10) * time.Second}}
	client.BaseURL, _ = url.Parse("http://api.paiza.io:80/runners/")

	return &client
}

// Request -
func (c *Client) Request(program *entity.Program) (*entity.ExecutionResult, error) {
	status, err := c.execProgram(program)
	if err != nil {
		return nil, errors.Wrap(err, "failed to exec program")
	}

	// wait execute program until completed
	for isCompleted := false; isCompleted == false; time.Sleep(1 * time.Second) {
		isCompleted, err = c.getStatus(status)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get status")
		}
	}

	return c.getResult(status)
}

// ExecProgramRequest is request to execute program
func (c *Client) execProgram(program *entity.Program) (*entity.Status, error) {
	query := map[string]string{"language": program.Lang, "api_key": "guest", "source_code": program.Program}
	values := url.Values{}
	for k, v := range query {
		values.Add(k, v)
	}

	urlString := strings.Join([]string{c.BaseURL.String(), string(create), values.Encode()}, "")
	req, err := http.NewRequest(http.MethodPost, urlString, nil)
	log.Printf("⚡️  %s", urlString)

	if err != nil {
		return nil, errors.Wrap(err, "failed to init paiza create request")
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to request paiza create request")
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var status entity.Status
	err = decoder.Decode(&status)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode paiza status")
	}

	return &status, nil
}

// GetStatusRequest is request to get execution status
func (c *Client) getStatus(status *entity.Status) (bool, error) {
	query := map[string]string{"id": status.ID, "api_key": "guest"}
	values := url.Values{}
	for k, v := range query {
		values.Add(k, v)
	}

	urlString := strings.Join([]string{c.BaseURL.String(), string(getStatus), values.Encode()}, "")

	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	if err != nil {
		return false, errors.Wrap(err, "failed to init paiza get_status request")
	}

	log.Printf("⚡️  %s", urlString)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return false, errors.Wrap(err, "failed to request paiza get_status request")
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(status)
	if err != nil {
		return false, errors.Wrap(err, "failed to decode paiza status")
	}

	return status.Status == "completed", nil
}

// GetResultRequest is request to get execution result
func (c *Client) getResult(status *entity.Status) (*entity.ExecutionResult, error) {
	query := map[string]string{"id": status.ID, "api_key": "guest"}
	values := url.Values{}
	for k, v := range query {
		values.Add(k, v)
	}

	urlString := strings.Join([]string{c.BaseURL.String(), string(getDetail), values.Encode()}, "")

	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init paiza get_details request")
	}

	log.Printf("⚡️  %s", urlString)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to request paiza get_details request")
	}
	defer resp.Body.Close()

	log.Println(resp.Body)

	decoder := json.NewDecoder(resp.Body)
	var result entity.ExecutionResult
	err = decoder.Decode(&result)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode paiza result")
	}

	return &result, nil
}
