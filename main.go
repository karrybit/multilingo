package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	paizaURL = "http://api.paiza.io:80/runners/"
	createPath = "create"
	detailPath = "get_details"
)

// TODO: naming
type CreateResponse struct {
	ID string `json:"id"`
	Status string `json:"status"`
}

func main() {
	values := url.Values{}
	values.Add("language", "swift")
	values.Add("api_key", "guest")
	values.Add("source_code", "print(114514)")

	resp, err := http.PostForm(paizaURL + createPath, values)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	// TODO: naming
	var cResp CreateResponse
	json.Unmarshal(bytes, &cResp)

	fmt.Println(cResp)
}
