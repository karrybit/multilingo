package main

import (
	"fmt"
	"time"
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

type DetailResponse struct {
	ID string `json:"id"`
	Language string `json:"language"`
	Note string `json:"note"`
	Status string `json:"status"`
	Build_stdout string `json:"build_stdout"`
	Build_stderr string `json:"build_stderr"`
	Build_exit_code int `json:"build_exit_code"`
	Build_time string `json:"build_time"`
	Build_memory int `json:"build_memory"`
	Build_result string `json:"build_result"`
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
	Exit_code int `json:"exit_code"`
	Time string `json:"time"`
	Memory int `json:"memory"`
	Connections int `json:"connections"`
	Result string `json:"result"`
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

	fmt.Printf("id: %s, status: %s\n", cResp.ID, cResp.Status)

	// TODO: goroutin使う
	time.Sleep(2 * time.Second)

	// TODO: naming
	values_2 := url.Values{}
	values_2.Add("id", cResp.ID)
	values_2.Add("api_key", "guest")
	resp_2, err := http.Get(paizaURL + detailPath + "?" + values_2.Encode())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp_2.Body.Close()

	fmt.Println(resp_2.Request.URL)

	// TODO: naming
	bytes_2, err := ioutil.ReadAll(resp_2.Body)
	if err != nil {
		fmt.Println(err)
	}

	// TODO: naming
	var dResp DetailResponse
	json.Unmarshal(bytes_2, &dResp)

	fmt.Printf("%v\n", dResp)
}
