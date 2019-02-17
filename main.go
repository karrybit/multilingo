package main

import (
	"fmt"
	"time"

	"github.com/TakumiKaribe/MultilinGo/logger"
	"github.com/TakumiKaribe/MultilinGo/model"
	"github.com/TakumiKaribe/MultilinGo/request"
)

func main() {
	id := execProgram()
	getResult(id)
	data := model.SlackRequestData{
		Token:    "gwbfPbBL1HEng7ZF0V8KOHAh",
		Channel:  "CG6LK3GSF",
		Text:     "114514",
		UserName: "UG7SDTG87",
	}
	data.Notification()
}

func execProgram() string {
	// TODO: language type
	query := map[string]string{"language": "swift", "api_key": "guest"}
	// TODO: add after parse
	query["source_code"] = "print(114514)"

	ch := make(chan request.StatusResult)
	go request.ExecProgramRequest(query, ch)

	result := <-ch

	if result.Err != nil {
		fmt.Println(result.Err)
		return ""
	}

	logger.PrintFields(&result.Response)

	return result.Response.ID
}

func getResult(id string) {
	query := map[string]string{"id": id, "api_key": "guest"}

	// wait execute program until completed
	for status := "runnig"; status != "completed"; time.Sleep(1 * time.Second) {
		ch := make(chan request.StatusResult)
		go request.GetStatusRequest(query, ch)
		statusResult := <-ch
		status = statusResult.Response.Status
	}

	ch := make(chan request.ExecutionResult)
	go request.GetResultRequest(query, ch)

	detailResult := <-ch

	if detailResult.Err != nil {
		fmt.Println(detailResult.Err)
		return
	}

	logger.PrintFields(&detailResult.Response)
}
