package request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/TakumiKaribe/MultilinGo/model"
)

type GetDetailsResult struct {
	Response model.DetailResponse
	Err      error
}

func GetDetails(query map[string]string, ch chan<- GetDetailsResult) {
	values := url.Values{}
	for k, v := range query {
		values.Add(k, v)
	}

	result := GetDetailsResult{}

	resp, err := http.Get(baseURL + detailPath + "?" + values.Encode())
	fmt.Printf("\n⚡️  %s\n\n", resp.Request.URL)

	if err != nil {
		result.Err = err
		ch <- result
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var detailResponse model.DetailResponse
	err = decoder.Decode(&detailResponse)
	if err != nil {
		result.Err = err
		ch <- result
	}

	result.Response = detailResponse
	ch <- result
}
