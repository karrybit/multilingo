package main

import (
	"fmt"
	"time"

	"github.com/TakumiKaribe/MultilinGo/request"
)

func main() {
	query := map[string]string{}
	query["language"] = "swift"
	// TODO: add after parse
	query["source_code"] = "print(114514)"
	createResponse, err := request.CreateID(query)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v\n", *createResponse)

	// TODO: goroutin使う
	time.Sleep(2 * time.Second)

	detailsResponse, err := request.GetDetails(map[string]string{"id": createResponse.ID})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v\n", *detailsResponse)
}
