package entity

import "fmt"

type message struct {
	status status
	output string
	time   string
	memory int
}

func (m *message) build() string {
	text := fmt.Sprintf("time: %s sec.\nmemory used: %d bytes", m.time, m.memory)
	if len(m.output) > 0 {
		text += fmt.Sprintf("\nlog:\n```%s```", m.output)
	}
	return text
}
