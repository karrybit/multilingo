package config

import (
	"fmt"

	"github.com/TakumiKaribe/multilingo/model"
	"github.com/kelseyhightower/envconfig"
)

// Config -
type Config struct {
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
}

// NewConfig -
func NewConfig() (*Config, error) {
	var config Config
	err := envconfig.Process("", &config) // env variable like MGO_AUTH_TOKEN
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// NewBotInfo -
func (c *Config) NewBotInfo(id string) (*model.Bot, error) {
	var config Config
	err := envconfig.Process("", &config) // env variable like MGO_AUTH_TOKEN
	if err != nil {
		return nil, err
	}

	switch id {
	case config.SwiftAppID:
		return &model.Bot{Name: "Swift", Token: config.SwiftOauthToken, Language: "swift"}, nil
	case config.PythonAppID:
		return &model.Bot{Name: "Python", Token: config.PythonOauthToken, Language: "python3"}, nil
	default:
		return nil, fmt.Errorf("No bot corresponding to %s was found", id)
	}
}
