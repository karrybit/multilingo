package buildMessage

import (
	"fmt"
	"strings"

	"github.com/TakumiKaribe/multilingo/entity/paiza"
	"github.com/TakumiKaribe/multilingo/entity/slack"
)

type status string
type color string
type title string

const (
	isSuccess status = "success"
	isFailure status = "failure"
	isError   status = "error"

	good    color = "good"
	warning color = "warning"
	danger  color = "danger"
)

type message struct {
	status status
	output string
	time   string
}

func (m *message) build() string {
	text := fmt.Sprintf("time: %s sec", m.time)
	if len(m.output) > 0 {
		text += fmt.Sprintf("\n```%s```", m.output)
	}
	return text
}

// MakeMessage -
func MakeMessage(result *paiza.Result) *[]*slack.Attachment {
	var attachments []*slack.Attachment

	if result.BuildResult == "" {
		attachments = append(attachments, makeExecResultAttachment(result))

	} else {
		attachments = *makeBuildResultAttachment(result)
	}

	return &attachments
}

func makeExecResultAttachment(result *paiza.Result) *slack.Attachment {
	attachment := slack.Attachment{Title: fmt.Sprintf("[EXEC %s]", strings.ToUpper(result.Result))}
	message := message{status: status(result.Result), time: result.Time}

	if message.status == isSuccess {
		message.output = result.Stdout
		attachment.Color = string(good)

	} else if message.status == isFailure {
		message.output = result.Stderr
		attachment.Color = string(warning)

	} else {
		message.output = result.Stderr
		attachment.Color = string(danger)
	}

	attachment.Text = message.build()
	return &attachment
}

func makeBuildResultAttachment(result *paiza.Result) *[]*slack.Attachment {
	var attachments []*slack.Attachment

	attachment := slack.Attachment{Title: fmt.Sprintf("[BUILD %s]", strings.ToUpper(result.BuildResult))}
	message := message{status: status(result.BuildResult), time: result.BuildTime}

	if message.status == isSuccess {
		message.output = result.BuildStdout
		attachment.Color = string(good)

	} else if message.status == isFailure {
		message.output = result.BuildStderr
		attachment.Color = string(warning)

	} else {
		message.output = result.BuildStderr
		attachment.Color = string(danger)
	}

	attachment.Text = message.build()
	attachments = append(attachments, &attachment)

	if message.status == isSuccess {
		attachments = append(attachments, makeExecResultAttachment(result))
	}

	return &attachments
}
