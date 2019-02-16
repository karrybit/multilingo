package request

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/TakumiKaribe/MultilinGo/model"
)

func CreateID(query map[string]string) (*model.CreateResponse, error) {
	values := url.Values{}
	for k, v := range query {
		values.Add(k, v)
	}

	resp, err := http.PostForm(baseURL+createPath, values)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var createResponse model.CreateResponse
	json.Unmarshal(bytes, &createResponse)

	return &createResponse, nil
}
