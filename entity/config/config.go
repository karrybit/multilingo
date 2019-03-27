package config

import (
	"fmt"

	"github.com/TakumiKaribe/multilingo/entity/slack"
	"github.com/kelseyhightower/envconfig"
)

// SharedConfig -
var SharedConfig *config

// Config -
type config struct {
	Debug bool `default:"true"`

	SlackPath string `required:"true" split_words:"true"`
	// App ID for each language
	// BashAppID   string `required:"true" split_words:"true"`
	// CppAppID    string `required:"true" split_words:"true"`
	PythonAppID string `required:"true" split_words:"true"`
	SwiftAppID  string `required:"true" split_words:"true"`

	// BotUserOAuthAccessToken for each language
	// BashOAuthToken   string `required:"true" split_words:"true"`
	// CppOAuthToken    string `required:"true" split_words:"true"`
	PythonOauthToken string `required:"true" split_words:"true"`
	SwiftOauthToken  string `required:"true" split_words:"true"`

	DebugConfig *debugConfig `required:"false"`
}

// DebugConfig -
type debugConfig struct {
	Channel                 string `required:"false"`
	PythonVerificationToken string `required:"false" split_words:"true"`
	SwiftVerificationToken  string `required:"false" split_words:"true"`
}

// Load -
func Load() error {
	if SharedConfig != nil {
		return nil
	}

	var _config config
	err := envconfig.Process("", &_config) // env variable like MGO_AUTH_TOKEN
	if err != nil {
		return err
	}

	if _config.Debug {
		var _debugConfig debugConfig
		err := envconfig.Process("", &_debugConfig) // env variable like MGO_AUTH_TOKEN
		if err != nil {
			return err
		}

		_config.DebugConfig = &_debugConfig
	}

	SharedConfig = &_config

	return nil
}

// NewBotInfo -
func (c *config) NewBotInfo(id string) (*slack.Bot, error) {
	switch id {
	case c.SwiftAppID:
		return &slack.Bot{Name: "Swift", Token: c.SwiftOauthToken, Language: "swift"}, nil
	case c.PythonAppID:
		return &slack.Bot{Name: "Python", Token: c.PythonOauthToken, Language: "python3"}, nil
	default:
		return nil, fmt.Errorf("No bot corresponding to %s was found", id)
	}
}
