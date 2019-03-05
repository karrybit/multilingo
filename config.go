package main

import "github.com/kelseyhightower/envconfig"

// Config -
type Config struct {
	Debug         bool `default:"false"`
	LogFormatJSON bool `default:"true"  split_words:"true"`
	// Authentication token for each language
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
	KotlinToken     string `required:"true" split_words:"true"`
}

func newConfig() (*Config, error) {
	var config Config
	err := envconfig.Process("", &config) // env variable like MGO_AUTH_TOKEN
	if err != nil {
		return nil, err
	}

	return &config, nil
}
