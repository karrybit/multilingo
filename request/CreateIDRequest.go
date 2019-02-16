package request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/TakumiKaribe/MultilinGo/model"
)

type CreateIDResult struct {
	Response model.CreateResponse
	Err      error
}

func CreateID(query map[string]string, ch chan<- CreateIDResult) {
	values := url.Values{}
	for k, v := range query {
		values.Add(k, v)
	}

	result := CreateIDResult{}

	resp, err := http.PostForm(baseURL+createPath, values)
	fmt.Printf("\n⚡️  %s\n\n", resp.Request.URL)
	if err != nil {
		result.Err = err
		ch <- result
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var createResponse model.CreateResponse
	err = decoder.Decode(&createResponse)
	if err != nil {
		result.Err = err
		ch <- result
	}

	result.Response = createResponse
	ch <- result
}
