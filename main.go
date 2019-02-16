package main

import (
	"fmt"

	"github.com/TakumiKaribe/MultilinGo/log"
	"github.com/TakumiKaribe/MultilinGo/request"
)

func main() {
	query := map[string]string{}
	query["language"] = "swift"
	// TODO: add after parse
	query["source_code"] = "print(114514)"
	query["api_key"] = "guest"

	// --- REQUEST ---
	createChannel := make(chan request.CreateIDResult)
	go request.CreateID(query, createChannel)

	createResult := <-createChannel

	// --- ERROR ---
	if createResult.Err != nil {
		fmt.Println(createResult.Err)
		return
	}

	log.PrintFields(&createResult.Response)

	// --- REQUEST ---
	detailChannel := make(chan request.GetDetailsResult)
	go request.GetDetails(map[string]string{"id": createResult.Response.ID, "api_key": "guest"}, detailChannel)

	detailResult := <-detailChannel

	// --- ERROR ---
	if detailResult.Err != nil {
		fmt.Println(detailResult.Err)
		return
	}

	log.PrintFields(&detailResult.Response)
}
