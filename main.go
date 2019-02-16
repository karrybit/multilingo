package main

import (
	"fmt"

	"github.com/TakumiKaribe/MultilinGo/Log"
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

	// --- LOG ---
	fmt.Println("⚙️  pre <-createChannel")
	createResult := <-createChannel
	fmt.Println("️⚙️  post <-createChannel")

	// --- ERROR ---
	if createResult.Err != nil {
		fmt.Println(createResult.Err)
		return
	}

	Log.PrintFields(&createResult.Response)

	// --- REQUEST ---
	detailChannel := make(chan request.GetDetailsResult)
	go request.GetDetails(map[string]string{"id": createResult.Response.ID, "api_key": "guest"}, detailChannel)

	// --- LOG ---
	fmt.Println("️⚙️  pre <-detailChannel")
	detailResult := <-detailChannel
	fmt.Println("⚙️  post <-detailChannel")

	// --- ERROR ---
	if detailResult.Err != nil {
		fmt.Println(detailResult.Err)
		return
	}

	Log.PrintFields(&detailResult.Response)
}
