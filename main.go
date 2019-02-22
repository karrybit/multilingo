package main

import (
	"time"

	"github.com/TakumiKaribe/MultilinGo/model"
	"github.com/TakumiKaribe/MultilinGo/parserawtext"
	"github.com/TakumiKaribe/MultilinGo/request"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Debug         bool `default:"false"`
	LogFormatJson bool `default:"true"  split_words:"true"`
	// Authentication token for each language
	CToken          string `required:"true" split_words:"true"`
	CppToken        string `required:"true" split_words:"true"`
	CsharpToken     string `required:"true" split_words:"true"`
	JavaToken       string `required:"true" split_words:"true"`
	Python3Token    string `required:"true" split_words:"true"`
	RubyToken       string `required:"true" split_words:"true"`
	JavascriptToken string `required:"true" split_words:"true"`
	ScalaToken      string `required:"true" split_words:"true"`
	GoToken         string `required:"true" split_words:"true"`
	HaskellToken    string `required:"true" split_words:"true"`
	RustToken       string `required:"true" split_words:"true"`
	SwiftToken      string `required:"true" split_words:"true"`
}

func main() {
	var conf Config
	err := envconfig.Process("mgo", &conf) // env variable like MGO_AUTH_TOKEN
	if err != nil {
		log.Fatal(err.Error())
	}
	// default level is INFO
	if conf.Debug {
		log.SetLevel(log.DebugLevel)
	}
	// default log format is ASCII
	if conf.LogFormatJson {
		log.SetFormatter(&log.JSONFormatter{})
	}

	// TODO: receive lambda context instead of string
	lambdaInput := "<@UG6LTEJBV>\n```print(114514)```\n"
	lang, text, err := parserawtext.Parse(lambdaInput)
	if err != nil {
		// TODO: response slack notification
		log.Fatal("failed to parse request. err: ", err)
	}
	status := execProgram(lang, text)
	getResult(status)
}

func execProgram(lang string, program string) model.Status {
	// TODO: language type
	query := map[string]string{"language": lang, "api_key": "guest", "source_code": program}

	ch := make(chan request.StatusResult)
	go request.ExecProgramRequest(query, ch)

	result := <-ch

	if result.Err != nil {
		log.Warn(result.Err)
		return model.Status{}
	}

	log.Debug(&result.Response)

	return result.Response
}

func getResult(status model.Status) {
	query := map[string]string{"id": status.ID, "api_key": "guest"}

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
		log.Warn(detailResult.Err)
		return
	}

	log.Debug(&detailResult.Response)
}
