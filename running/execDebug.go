package running

import (
	"flag"
	"os"

	"github.com/TakumiKaribe/multilingo/config"
	"github.com/TakumiKaribe/multilingo/model"
	log "github.com/sirupsen/logrus"
)

// ExecDebug -
func ExecDebug() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.JSONFormatter{})

	// setup config
	config, err := config.NewConfig()
	if err != nil {
		log.Warn(err.Error())
	}

	// init dummy
	flag.Parse()
	event := model.Event{Text: "<@debug>```" + flag.Arg(0) + "```", Channel: os.Getenv("CHANNEL")}
	requestBody := model.APIGateWayRequestBody{Token: os.Getenv("SWIFT_VERIFICATION_TOKEN"), APIAppID: config.SwiftAppID, Event: event}

	run(&requestBody)
}
