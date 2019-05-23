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
	ObjcAppID         string `required:"false" split_words:"true"`
	JavaAppID         string `required:"false" split_words:"true"`
	KotlinAppID       string `required:"false" split_words:"true"`
	ScalaAppID        string `required:"false" split_words:"true"`
	SwiftAppID        string `required:"false" split_words:"true"`
	CSharpAppID       string `required:"false" split_words:"true"`
	HaskellAppID      string `required:"false" split_words:"true"`
	GoAppID           string `required:"false" split_words:"true"`
	ErlangAppID       string `required:"false" split_words:"true"`
	PerlAppID         string `required:"false" split_words:"true"`
	PythonAppID       string `required:"false" split_words:"true"`
	RubyAppID         string `required:"false" split_words:"true"`
	PhpAppID          string `required:"false" split_words:"true"`
	BashAppID         string `required:"false" split_words:"true"`
	RAppID            string `required:"false" split_words:"true"`
	JavaScriptAppID   string `required:"false" split_words:"true"`
	CoffeeScriptAppID string `required:"false" split_words:"true"`
	VbAppID           string `required:"false" split_words:"true"`
	CobolAppID        string `required:"false" split_words:"true"`
	FSharpAppID       string `required:"false" split_words:"true"`
	DAppID            string `required:"false" split_words:"true"`
	ClojureAppID      string `required:"false" split_words:"true"`
	ElixierAppID      string `required:"false" split_words:"true"`
	MysqlAppID        string `required:"false" split_words:"true"`
	RustAppID         string `required:"false" split_words:"true"`
	SchemeAppID       string `required:"false" split_words:"true"`
	CommonLispAppID   string `required:"false" split_words:"true"`

	// BotUserOAuthAccessToken for each language
	COauthToken            string `required:"false" split_words:"true"`
	CppOauthToken          string `required:"false" split_words:"true"`
	ObjcOauthToken         string `required:"false" split_words:"true"`
	JavaOauthToken         string `required:"false" split_words:"true"`
	KotlinOauthToken       string `required:"false" split_words:"true"`
	ScalaOauthToken        string `required:"false" split_words:"true"`
	SwiftOauthToken        string `required:"false" split_words:"true"`
	CSharpOauthToken       string `required:"false" split_words:"true"`
	HaskellOauthToken      string `required:"false" split_words:"true"`
	GoOauthToken           string `required:"false" split_words:"true"`
	ErlangOauthToken       string `required:"false" split_words:"true"`
	PerlOauthToken         string `required:"false" split_words:"true"`
	PythonOauthToken       string `required:"false" split_words:"true"`
	RubyOauthToken         string `required:"false" split_words:"true"`
	PhpOauthToken          string `required:"false" split_words:"true"`
	BashOauthToken         string `required:"false" split_words:"true"`
	ROauthToken            string `required:"false" split_words:"true"`
	JavaScriptOauthToken   string `required:"false" split_words:"true"`
	CoffeeScriptOauthToken string `required:"false" split_words:"true"`
	VbOauthToken           string `required:"false" split_words:"true"`
	CobolOauthToken        string `required:"false" split_words:"true"`
	FSharpOauthToken       string `required:"false" split_words:"true"`
	DOauthToken            string `required:"false" split_words:"true"`
	ClojureOauthToken      string `required:"false" split_words:"true"`
	ElixierOauthToken      string `required:"false" split_words:"true"`
	MysqlOauthToken        string `required:"false" split_words:"true"`
	RustOauthToken         string `required:"false" split_words:"true"`
	SchemeOauthToken       string `required:"false" split_words:"true"`
	CommonLispOauthToken   string `required:"false" split_words:"true"`

	DebugConfig *debugConfig `required:"false"`
}

// DebugConfig -
type debugConfig struct {
	Channel string `required:"false"`

	// BotVerificationToken for each language
	CVerificationToken            string `required:"false" split_words:"true"`
	CppVerificationToken          string `required:"false" split_words:"true"`
	ObjcVerificationToken         string `required:"false" split_words:"true"`
	JavaVerificationToken         string `required:"false" split_words:"true"`
	KotlinVerificationToken       string `required:"false" split_words:"true"`
	ScalaVerificationToken        string `required:"false" split_words:"true"`
	SwiftVerificationToken        string `required:"false" split_words:"true"`
	CSharpVerificationToken       string `required:"false" split_words:"true"`
	GoVerificationToken           string `required:"false" split_words:"true"`
	HaskellVerificationToken      string `required:"false" split_words:"true"`
	ErlangVerificationToken       string `required:"false" split_words:"true"`
	PerlVerificationToken         string `required:"false" split_words:"true"`
	PythonVerificationToken       string `required:"false" split_words:"true"`
	RubyVerificationToken         string `required:"false" split_words:"true"`
	PhpVerificationToken          string `required:"false" split_words:"true"`
	BashVerificationToken         string `required:"false" split_words:"true"`
	RVerificationToken            string `required:"false" split_words:"true"`
	JavaScriptVerificationToken   string `required:"false" split_words:"true"`
	CoffeeScriptVerificationToken string `required:"false" split_words:"true"`
	VbVerificationToken           string `required:"false" split_words:"true"`
	CobolVerificationToken        string `required:"false" split_words:"true"`
	FSharpVerificationToken       string `required:"false" split_words:"true"`
	DVerificationToken            string `required:"false" split_words:"true"`
	ClojureVerificationToken      string `required:"false" split_words:"true"`
	ElixierVerificationToken      string `required:"false" split_words:"true"`
	MysqlVerificationToken        string `required:"false" split_words:"true"`
	RustVerificationToken         string `required:"false" split_words:"true"`
	SchemeVerificationToken       string `required:"false" split_words:"true"`
	CommonLispVerificationToken   string `required:"false" split_words:"true"`
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
	case c.CAppID:
		return &slack.Bot{Name: "C", Token: c.COauthToken, Language: "c"}, nil
	case c.CppAppID:
		return &slack.Bot{Name: "Cpp", Token: c.CppOauthToken, Language: "cpp"}, nil
	case c.ObjcAppID:
		return &slack.Bot{Name: "Objective-C", Token: c.ObjcOauthToken, Language: "objective-c"}, nil
	case c.JavaAppID:
		return &slack.Bot{Name: "Java", Token: c.JavaOauthToken, Language: "java"}, nil
	case c.KotlinAppID:
		return &slack.Bot{Name: "Kotlin", Token: c.KotlinOauthToken, Language: "kotlin"}, nil
	case c.ScalaAppID:
		return &slack.Bot{Name: "Scala", Token: c.KotlinOauthToken, Language: "scala"}, nil
	case c.SwiftAppID:
		return &slack.Bot{Name: "Swift", Token: c.SwiftOauthToken, Language: "swift"}, nil
	case c.CSharpAppID:
		return &slack.Bot{Name: "Csharp", Token: c.CSharpOauthToken, Language: "csharp"}, nil
	case c.GoAppID:
		return &slack.Bot{Name: "Go", Token: c.GoOauthToken, Language: "go"}, nil
	case c.HaskellAppID:
		return &slack.Bot{Name: "Haskell", Token: c.HaskellOauthToken, Language: "haskell"}, nil
	case c.ErlangAppID:
		return &slack.Bot{Name: "Erlang", Token: c.HaskellOauthToken, Language: "erlang"}, nil
	case c.PerlAppID:
		return &slack.Bot{Name: "Perl", Token: c.HaskellOauthToken, Language: "perl"}, nil
	case c.PythonAppID:
		return &slack.Bot{Name: "Python", Token: c.PythonOauthToken, Language: "python3"}, nil
	case c.RubyAppID:
		return &slack.Bot{Name: "Ruby", Token: c.PythonOauthToken, Language: "ruby"}, nil
	case c.PhpAppID:
		return &slack.Bot{Name: "PHP", Token: c.PythonOauthToken, Language: "php"}, nil
	case c.BashAppID:
		return &slack.Bot{Name: "Bash", Token: c.PythonOauthToken, Language: "bash"}, nil
	case c.RAppID:
		return &slack.Bot{Name: "R", Token: c.PythonOauthToken, Language: "r"}, nil
	case c.JavaScriptAppID:
		return &slack.Bot{Name: "JavaScript", Token: c.PythonOauthToken, Language: "javascript"}, nil
	case c.CoffeeScriptAppID:
		return &slack.Bot{Name: "CoffeeScript", Token: c.PythonOauthToken, Language: "coffeescript"}, nil
	case c.VbAppID:
		return &slack.Bot{Name: "VB", Token: c.PythonOauthToken, Language: "vb"}, nil
	case c.CobolAppID:
		return &slack.Bot{Name: "COBOL", Token: c.PythonOauthToken, Language: "cobol"}, nil
	case c.FSharpAppID:
		return &slack.Bot{Name: "Fsharp", Token: c.PythonOauthToken, Language: "fsharp"}, nil
	case c.DAppID:
		return &slack.Bot{Name: "D", Token: c.PythonOauthToken, Language: "d"}, nil
	case c.ClojureAppID:
		return &slack.Bot{Name: "Clojure", Token: c.PythonOauthToken, Language: "clojure"}, nil
	case c.ElixierAppID:
		return &slack.Bot{Name: "Elixier", Token: c.PythonOauthToken, Language: "elixier"}, nil
	case c.MysqlAppID:
		return &slack.Bot{Name: "MySQL", Token: c.PythonOauthToken, Language: "mysql"}, nil
	case c.RustAppID:
		return &slack.Bot{Name: "Rust", Token: c.PythonOauthToken, Language: "rust"}, nil
	case c.SchemeAppID:
		return &slack.Bot{Name: "Scheme", Token: c.PythonOauthToken, Language: "scheme"}, nil
	case c.CommonLispAppID:
		return &slack.Bot{Name: "CommonLisp", Token: c.PythonOauthToken, Language: "commonlisp"}, nil

	default:
		return nil, multilingoerror.New(multilingoerror.NewBotCorrespondingToID, id, "")
	}
}
