package parserawtext

import (
	"regexp"

	"github.com/pkg/errors"
)

var regex = regexp.MustCompile(`(?s)^.*<@.*>.*\x60\x60\x60(.*)\x60\x60\x60.*$`)

// TODO: receive lambda context instead of string
func Parse(text string) (program string, err error) {
	param := regex.FindStringSubmatch(text)
	if len(param) != 2 {
		err = errors.New("failed to parse post messages.")
		return
	}
	program = param[1]
	return
}
