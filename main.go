package main

import (
	"fmt"
	"time"

	"github.com/TakumiKaribe/MultilinGo/log"
	"github.com/TakumiKaribe/MultilinGo/request"
)

func main() {
	id := f()
	g(id)
}

func f() string {
	// --- SETUP ---
	query := map[string]string{}
	query["language"] = "swift"
	// TODO: add after parse
	query["source_code"] = "print(114514)"
	query["api_key"] = "guest"

	// --- REQUEST ---
	ch := make(chan request.StatusResult)
	go request.ExecProgramRequest(query, ch)

	result := <-ch

	// --- ERROR ---
	if result.Err != nil {
		fmt.Println(result.Err)
		return ""
	}

	// --- LOG ---
	log.PrintFields(&result.Response)

	return result.Response.ID
}

func g(id string) {
	for status := "runnig"; status != "completed"; time.Sleep(1 * time.Second) {
		ch := make(chan request.StatusResult)
		go request.GetStatusRequest(map[string]string{"id": id, "api_key": "guest"}, ch)
		statusResult := <-ch
		status = statusResult.Response.Status
	}

	// --- REQUEST ---
	ch := make(chan request.ExecutionResult)
	go request.GetResultRequest(map[string]string{"id": id, "api_key": "guest"}, ch)

	detailResult := <-ch

	// --- ERROR ---
	if detailResult.Err != nil {
		fmt.Println(detailResult.Err)
		return
	}

	// --- LOG ---
	log.PrintFields(&detailResult.Response)
}
