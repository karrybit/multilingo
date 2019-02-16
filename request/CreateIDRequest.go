package request

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	var createResponse model.CreateResponse
	json.Unmarshal(bytes, &createResponse)

	fmt.Printf("📦  %v\n", createResponse)

	result.Response = createResponse
	fmt.Println("⚙️  ch <- CreateIDResult")
	ch <- result
}
