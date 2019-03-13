package running

import (
	"flag"
	"fmt"
	"os"

	"github.com/TakumiKaribe/multilingo/model"
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
	event := model.Event{Text: "<@debug>```" + flag.Arg(0) + "```", Channel: os.Getenv("CHANNEL")}
	requestBody := model.APIGateWayRequestBody{Token: token, APIAppID: appID, Event: event}

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
		return "", "", fmt.Errorf("No config corresponding to %s was found", arg)
	}
}
