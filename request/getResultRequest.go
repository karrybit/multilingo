package request

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/TakumiKaribe/MultilinGo/model"
)

// ExecutionResult -
type ExecutionResult struct {
	Response model.ExecutionResult
	Err      error
}

// GetResultRequest is request to get execution result
func GetResultRequest(query map[string]string, ch chan<- ExecutionResult) {
	values := url.Values{}
	for k, v := range query {
		values.Add(k, v)
	}

	result := ExecutionResult{}

	resp, err := http.Get(baseURL + detailPath + "?" + values.Encode())
	log.Printf("⚡️  %s\n", resp.Request.URL)

	if err != nil {
		result.Err = err
		ch <- result
	}

	defer resp.Body.Close()

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
