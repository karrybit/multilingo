package request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/TakumiKaribe/MultilinGo/model"
)

// GetStatusRequest is request to get execution status
func GetStatusRequest(query map[string]string, ch chan<- StatusResult) {
	values := url.Values{}
	for k, v := range query {
		values.Add(k, v)
	}

	result := StatusResult{}

	resp, err := http.Get(baseURL + statusPath + "?" + values.Encode())
	fmt.Printf("\n⚡️  %s\n\n", resp.Request.URL)
	if err != nil {
		result.Err = err
		ch <- result
	}

	defer resp.Body.Close()

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
