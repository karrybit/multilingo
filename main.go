package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	values := url.Values{}
	values.Add("language", "swift")
	values.Add("api_key", "guest")
	values.Add("source_code", "print(114514)")

	resp, err := http.PostForm("http://api.paiza.io:80/runners/create", values)
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
