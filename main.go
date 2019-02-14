package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	paizaURL = "http://api.paiza.io:80/runners/"
	createPath = "create"
	detailPath = "get_details"
)

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

	b, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		fmt.Println(string(b))
	}
}
