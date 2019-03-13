package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// Config -
type Config struct {
	Debug         bool `default:"false"`
	LogFormatJSON bool `default:"true"  split_words:"true"`

	// App ID for each language
	// BashAppID   string `required:"true" split_words:"true"`
	// CppAppID    string `required:"true" split_words:"true"`
	PythonAppID string `required:"true" split_words:"true"`
	SwiftAppID  string `required:"true" split_words:"true"`

	// BotUserOAuthAccessToken for each language
	// BashOAuthToken   string `required:"true" split_words:"true"`
	// CppOAuthToken    string `required:"true" split_words:"true"`
	PythonOAuthToken string `required:"true" split_words:"true"`
	SwiftOAuthToken  string `required:"true" split_words:"true"`
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

// LookUpToken -
func (c *Config) LookUpToken(id string) (string, error) {
	switch id {
	case c.SwiftAppID:
		return c.SwiftOAuthToken, nil
	case c.PythonAppID:
		return c.PythonOAuthToken, nil
	default:
		return "", fmt.Errorf("No language corresponding to %s was found", id)
	}
}

// LookUpLanguage -
func (c *Config) LookUpLanguage(id string) (string, error) {
	switch id {
	case c.SwiftAppID:
		return "swift", nil
	case c.PythonAppID:
		return "python3", nil
	default:
		return "", fmt.Errorf("No language corresponding to %s was found", id)
	}
}
