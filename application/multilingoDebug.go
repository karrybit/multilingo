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
	case "c":
		return config.SharedConfig.DebugConfig.CVerificationToken, config.SharedConfig.CAppID, nil
	case "cpp":
		return config.SharedConfig.DebugConfig.CppVerificationToken, config.SharedConfig.CppAppID, nil
	case "objective-c":
		return config.SharedConfig.DebugConfig.ObjcVerificationToken, config.SharedConfig.ObjcAppID, nil
	case "java":
		return config.SharedConfig.DebugConfig.JavaVerificationToken, config.SharedConfig.JavaAppID, nil
	case "kotlin":
		return config.SharedConfig.DebugConfig.KotlinVerificationToken, config.SharedConfig.KotlinAppID, nil
	case "scala":
		return config.SharedConfig.DebugConfig.ScalaVerificationToken, config.SharedConfig.ScalaAppID, nil
	case "swift":
		return config.SharedConfig.DebugConfig.SwiftVerificationToken, config.SharedConfig.SwiftAppID, nil
	case "csharp":
		return config.SharedConfig.DebugConfig.CSharpVerificationToken, config.SharedConfig.CSharpAppID, nil
	case "go":
		return config.SharedConfig.DebugConfig.GoVerificationToken, config.SharedConfig.GoAppID, nil
	case "haskell":
		return config.SharedConfig.DebugConfig.HaskellVerificationToken, config.SharedConfig.HaskellAppID, nil
	case "erlang":
		return config.SharedConfig.DebugConfig.ErlangVerificationToken, config.SharedConfig.ErlangAppID, nil
	case "perl":
		return config.SharedConfig.DebugConfig.PerlVerificationToken, config.SharedConfig.PerlAppID, nil
	case "python":
		return config.SharedConfig.DebugConfig.PythonVerificationToken, config.SharedConfig.PythonAppID, nil
	case "ruby":
		return config.SharedConfig.DebugConfig.RubyVerificationToken, config.SharedConfig.RubyAppID, nil
	case "php":
		return config.SharedConfig.DebugConfig.PhpVerificationToken, config.SharedConfig.PhpAppID, nil
	case "bash":
		return config.SharedConfig.DebugConfig.BashVerificationToken, config.SharedConfig.BashAppID, nil
	case "r":
		return config.SharedConfig.DebugConfig.RVerificationToken, config.SharedConfig.RAppID, nil
	case "javascript":
		return config.SharedConfig.DebugConfig.JavaScriptVerificationToken, config.SharedConfig.JavaScriptAppID, nil
	case "coffeeScript":
		return config.SharedConfig.DebugConfig.CoffeeScriptVerificationToken, config.SharedConfig.CoffeeScriptAppID, nil
	case "vb":
		return config.SharedConfig.DebugConfig.VbVerificationToken, config.SharedConfig.VbAppID, nil
	case "cobol":
		return config.SharedConfig.DebugConfig.CobolVerificationToken, config.SharedConfig.CobolAppID, nil
	case "fsharp":
		return config.SharedConfig.DebugConfig.FSharpVerificationToken, config.SharedConfig.FSharpAppID, nil
	case "d":
		return config.SharedConfig.DebugConfig.DVerificationToken, config.SharedConfig.DAppID, nil
	case "clojure":
		return config.SharedConfig.DebugConfig.ClojureVerificationToken, config.SharedConfig.ClojureAppID, nil
	case "elixier":
		return config.SharedConfig.DebugConfig.ElixierVerificationToken, config.SharedConfig.ElixierAppID, nil
	case "mysql":
		return config.SharedConfig.DebugConfig.MysqlVerificationToken, config.SharedConfig.MysqlAppID, nil
	case "rust":
		return config.SharedConfig.DebugConfig.RustVerificationToken, config.SharedConfig.RustAppID, nil
	case "scheme":
		return config.SharedConfig.DebugConfig.SchemeVerificationToken, config.SharedConfig.SchemeAppID, nil
	case "commonlisp":
		return config.SharedConfig.DebugConfig.CommonLispVerificationToken, config.SharedConfig.CommonLispAppID, nil
	default:
		return "", "", multilingoerror.New(multilingoerror.NotFoundConfig, arg, "")
	}
}
