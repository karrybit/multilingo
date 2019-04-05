package parsetext

import (
	"html"
	"regexp"

	"github.com/TakumiKaribe/multilingo/entity/multilingoerror"
)

var regex = regexp.MustCompile(`(?s)^.*<@.*>.*\x60\x60\x60(.*)\x60\x60\x60.*$`)

// Parse -
func Parse(text string) (program string, err error) {
	param := regex.FindStringSubmatch(text)
	if len(param) != 2 {
		err = multilingoerror.New(multilingoerror.ParseProgram, "", "")
		return
	}
	program = html.UnescapeString(param[1])
	return
}
