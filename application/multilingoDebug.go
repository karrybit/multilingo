package application

import (
	"flag"
	"os"

	"github.com/TakumiKaribe/multilingo/entity"
	"github.com/TakumiKaribe/multilingo/entity/config"
	"github.com/TakumiKaribe/multilingo/entity/multilingoerror"
	"github.com/TakumiKaribe/multilingo/logger"
)

// ExecDebug -
func ExecDebug() {
	// init dummy
	flag.Parse()
	token, appID, err := switchLanguage()
	if err != nil {
		logger.Log.Fatal(err)
	}
	event := entity.Event{Text: flag.Arg(1), Channel: os.Getenv("CHANNEL")}
	requestBody := entity.APIGateWayRequestBody{Token: token, APIAppID: appID, Event: event}

	run(&requestBody)
}

func switchLanguage() (string, string, error) {
	arg := flag.Arg(0)
	switch arg {
	case "swift":
		return config.SharedConfig.DebugConfig.SwiftVerificationToken, config.SharedConfig.SwiftAppID, nil
	case "python":
		return config.SharedConfig.DebugConfig.PythonVerificationToken, config.SharedConfig.PythonAppID, nil
	default:
		return "", "", multilingoerror.New(multilingoerror.NotFoundConfig, arg, "")
	}
}
