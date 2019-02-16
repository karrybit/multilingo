package request

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/TakumiKaribe/MultilinGo/model"
)

func GetDetails(query map[string]string) (*model.DetailResponse, error) {
	values := url.Values{}
	for k, v := range query {
		values.Add(k, v)
	}

	resp, err := http.Get(baseURL + detailPath + "?" + values.Encode())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var detailResponse model.DetailResponse
	json.Unmarshal(bytes, &detailResponse)

	return &detailResponse, nil
}
