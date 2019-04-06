package parsetext

import (
	"regexp"
	"strings"

	"github.com/TakumiKaribe/multilingo/entity/multilingoerror"
)

var (
	regex  = regexp.MustCompile(`.+?(?:\x60\x60\x60(.*?)\x60\x60\x60){1,2}`)
	regex2 = regexp.MustCompile(`\x60\x60\x60(.*?)\x60\x60\x60`)

	replace  = strings.NewReplacer("\n", "\\n", "\r", "\\n", "\r\n", "\\n")
	replace2 = strings.NewReplacer("\\n", "\n")
)

// Parse -
func Parse(text string) (input string, program string, err error) {
	text = replace.Replace(text)

	var str []string

	for _, match := range regex.FindAllString(text, -1) {
		for _, match2 := range regex2.FindStringSubmatch(match) {
			match2 = replace2.Replace(match2)
			str = append(str, match2)
		}
	}

	if len(str) == 2 {
		program = str[1]

	} else if len(str) == 4 {
		input = str[1]
		program = str[3]

	} else {
		err = multilingoerror.New(multilingoerror.ParseProgram, "", "")
	}

	return
}
