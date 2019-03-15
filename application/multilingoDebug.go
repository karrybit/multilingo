package application

import (
	"flag"
	"os"

	"github.com/TakumiKaribe/multilingo/entity"
	"github.com/TakumiKaribe/multilingo/entity/multilingoerror"
	log "github.com/sirupsen/logrus"
)

// ExecDebug -
func ExecDebug() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.JSONFormatter{})

	// init dummy
	flag.Parse()
	token, appID, err := switchLanguage()
	if err != nil {
		log.Fatal(err)
	}
	event := entity.Event{Text: "<@debug>```" + flag.Arg(0) + "```", Channel: os.Getenv("CHANNEL")}
	requestBody := entity.APIGateWayRequestBody{Token: token, APIAppID: appID, Event: event}

	run(&requestBody)
}

func switchLanguage() (string, string, error) {
	arg := flag.Arg(1)
	switch arg {
	case "swift":
		return os.Getenv("SWIFT_VERIFICATION_TOKEN"), os.Getenv("SWIFT_APP_ID"), nil
	case "python":
		return os.Getenv("PYTHON_VERIFICATION_TOKEN"), os.Getenv("PYTHON_APP_ID"), nil
	default:
		return "", "", multilingoerror.New(multilingoerror.NotFoundConfig, arg, "")
	}
}
