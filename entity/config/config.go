package config

import (
	"os"
	"strconv"

	"github.com/TakumiKaribe/multilingo/entity/multilingoerror"
	"github.com/TakumiKaribe/multilingo/entity/slack"
	"github.com/TakumiKaribe/multilingo/logger"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// SharedConfig -
var SharedConfig *config

// Config -
type config struct {
	Debug bool `default:"true"`

	SlackPath string `required:"true" split_words:"true"`

	// App ID for each language
	CAppID            string `required:"false" split_words:"true"`
	CppAppID          string `required:"false" split_words:"true"`
	CSharpAppID       string `required:"false" split_words:"true"`
	JavaAppID         string `required:"false" split_words:"true"`
	PythonAppID       string `required:"false" split_words:"true"`
	RubyAppID         string `required:"false" split_words:"true"`
	PerlAppID         string `required:"false" split_words:"true"`
	PhpAppID          string `required:"false" split_words:"true"`
	JavaScriptAppID   string `required:"false" split_words:"true"`
	ObjcAppID         string `required:"false" split_words:"true"`
	KotlinAppID       string `required:"false" split_words:"true"`
	ScalaAppID        string `required:"false" split_words:"true"`
	SwiftAppID        string `required:"false" split_words:"true"`
	GoAppID           string `required:"false" split_words:"true"`
	HaskellAppID      string `required:"false" split_words:"true"`
	CoffeeScriptAppID string `required:"false" split_words:"true"`
	BashAppID         string `required:"false" split_words:"true"`
	ErlangAppID       string `required:"false" split_words:"true"`
	RAppID            string `required:"false" split_words:"true"`
	CobolAppID        string `required:"false" split_words:"true"`
	VbAppID           string `required:"false" split_words:"true"`
	FSharpAppID       string `required:"false" split_words:"true"`
	ClojureAppID      string `required:"false" split_words:"true"`
	DAppID            string `required:"false" split_words:"true"`
	ElixierAppID      string `required:"false" split_words:"true"`
	RustAppID         string `required:"false" split_words:"true"`
	SchemeAppID       string `required:"false" split_words:"true"`
	MysqlAppID        string `required:"false" split_words:"true"`

	// BotUserOAuthAccessToken for each language
	COauthToken            string `required:"false" split_words:"true"`
	CppOauthToken          string `required:"false" split_words:"true"`
	CSharpOauthToken       string `required:"false" split_words:"true"`
	JavaOauthToken         string `required:"false" split_words:"true"`
	PythonOauthToken       string `required:"false" split_words:"true"`
	RubyOauthToken         string `required:"false" split_words:"true"`
	PerlOauthToken         string `required:"false" split_words:"true"`
	PhpOauthToken          string `required:"false" split_words:"true"`
	JavaScriptOauthToken   string `required:"false" split_words:"true"`
	ObjcOauthToken         string `required:"false" split_words:"true"`
	KotlinOauthToken       string `required:"false" split_words:"true"`
	ScalaOauthToken        string `required:"false" split_words:"true"`
	SwiftOauthToken        string `required:"false" split_words:"true"`
	GoOauthToken           string `required:"false" split_words:"true"`
	HaskellOauthToken      string `required:"false" split_words:"true"`
	CoffeeScriptOauthToken string `required:"false" split_words:"true"`
	BashOauthToken         string `required:"false" split_words:"true"`
	ErlangOauthToken       string `required:"false" split_words:"true"`
	ROauthToken            string `required:"false" split_words:"true"`
	CobolOauthToken        string `required:"false" split_words:"true"`
	VbOauthToken           string `required:"false" split_words:"true"`
	FSharpOauthToken       string `required:"false" split_words:"true"`
	ClojureOauthToken      string `required:"false" split_words:"true"`
	DOauthToken            string `required:"false" split_words:"true"`
	ElixierOauthToken      string `required:"false" split_words:"true"`
	RustOauthToken         string `required:"false" split_words:"true"`
	SchemeOauthToken       string `required:"false" split_words:"true"`
	MysqlOauthToken        string `required:"false" split_words:"true"`

	DebugConfig *debugConfig `required:"false"`
}

// DebugConfig -
type debugConfig struct {
	Channel string `required:"false"`

	// BotVerificationToken for each language
	CVerificationToken            string `required:"false" split_words:"true"`
	CppVerificationToken          string `required:"false" split_words:"true"`
	CSharpVerificationToken       string `required:"false" split_words:"true"`
	JavaVerificationToken         string `required:"false" split_words:"true"`
	PythonVerificationToken       string `required:"false" split_words:"true"`
	RubyVerificationToken         string `required:"false" split_words:"true"`
	PerlVerificationToken         string `required:"false" split_words:"true"`
	PhpVerificationToken          string `required:"false" split_words:"true"`
	JavaScriptVerificationToken   string `required:"false" split_words:"true"`
	ObjcVerificationToken         string `required:"false" split_words:"true"`
	KotlinVerificationToken       string `required:"false" split_words:"true"`
	ScalaVerificationToken        string `required:"false" split_words:"true"`
	SwiftVerificationToken        string `required:"false" split_words:"true"`
	GoVerificationToken           string `required:"false" split_words:"true"`
	HaskellVerificationToken      string `required:"false" split_words:"true"`
	CoffeeScriptVerificationToken string `required:"false" split_words:"true"`
	BashVerificationToken         string `required:"false" split_words:"true"`
	ErlangVerificationToken       string `required:"false" split_words:"true"`
	RVerificationToken            string `required:"false" split_words:"true"`
	CobolVerificationToken        string `required:"false" split_words:"true"`
	VbVerificationToken           string `required:"false" split_words:"true"`
	FSharpVerificationToken       string `required:"false" split_words:"true"`
	ClojureVerificationToken      string `required:"false" split_words:"true"`
	DVerificationToken            string `required:"false" split_words:"true"`
	ElixierVerificationToken      string `required:"false" split_words:"true"`
	RustVerificationToken         string `required:"false" split_words:"true"`
	SchemeVerificationToken       string `required:"false" split_words:"true"`
	MysqlVerificationToken        string `required:"false" split_words:"true"`
}

// Load -
func Load() error {
	if SharedConfig != nil {
		return nil
	}

	if b, _ := strconv.ParseBool(os.Getenv("DEBUG")); b {
		err := godotenv.Load()
		if err != nil {
			logger.Log.Fatal("Error loading .env file")
		}
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
		return nil, multilingoerror.New(multilingoerror.NewBotCorrespondingToID, id, "")
	}
}
