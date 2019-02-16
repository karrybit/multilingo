package request

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	fmt.Printf("⚡️  %s\n", resp.Request.URL)

	if err != nil {
		result.Err = err
		ch <- result
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		result.Err = err
		ch <- result
	}

	var detailResponse model.DetailResponse
	json.Unmarshal(bytes, &detailResponse)

	fmt.Printf("📦  %v\n", detailResponse)

	result.Response = detailResponse
	fmt.Println("⚙️  ch <- GetDetailsResult")
	ch <- result
}
