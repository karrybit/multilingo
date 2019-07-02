package parsetext

import (
	"html"
	"regexp"
	"strings"

	"multilingo/entity/multilingoerror"
	"multilingo/logger"
)

var (
	regex         = regexp.MustCompile(`(?s)(?:\x60\x60\x60(.*?)\x60\x60\x60){1,2}`)
	regex2        = regexp.MustCompile(`(?s)\x60\x60\x60(.*)\x60\x60\x60{1,2}`)
	regexReplaceN = regexp.MustCompile(`\\(\\*)n`)
	regexReplaceT = regexp.MustCompile(`\\(\\*)t`)
)

// Parse -
func Parse(text string) (input string, program string, err error) {
	var str []string

	logger.Log.Infof("text: %s\n", text)
	for _, match := range regex.FindAllString(text, -1) {
		for _, match2 := range regex2.FindStringSubmatch(match) {
			match2 = regexReplaceN.ReplaceAllString(match2, "\n")
			match2 = regexReplaceT.ReplaceAllString(match2, "\t")
			match2 = strings.Trim(match2, "\n")
			str = append(str, match2)
		}
	}

	if len(str) == 2 {
		program = html.UnescapeString(str[1])

	} else if len(str) == 4 {
		input = html.UnescapeString(str[1])
		program = html.UnescapeString(str[3])

	} else {
		err = multilingoerror.New(multilingoerror.ParseProgram, "", "")
	}

	return
}
