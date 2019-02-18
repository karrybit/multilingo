package parseRawText

import (
	"strings"
)

// TODO: receive lambda context instead of string
func Parse(text string) (lang string, program string, err error) {
	lang, err = parseLanguage(text)
	if err != err {
		return
	}

	return
}

func parseLanguage(text string) (lang string, err error) {
	// find first '@'
	findAtMark := func(c rune) bool { return c == '@' }
	// find first '>'
	findAtGreaterThan := func(c rune) bool { return c == '>' }
	id := text[strings.IndexFunc(text, findAtMark)+1 : strings.IndexFunc(text, findAtGreaterThan)]
	lang, err = lookUp(id)
	return
}
