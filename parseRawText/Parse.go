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

	program, err = parseProgram(text)
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

func parseProgram(text string) (program string, err error) {
	// find first '>'
	findAtGreaterThan := func(c rune) bool { return c == '>' }
	text = text[strings.IndexFunc(text, findAtGreaterThan)+1:]
	text = strings.TrimSpace(text)
	text = strings.Trim(text, "\n`")
	return text, err
}
