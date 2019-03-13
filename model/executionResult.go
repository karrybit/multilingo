package model

import (
	"fmt"
	"strings"
)

// ExecutionResult -
type ExecutionResult struct {
	ID            string `json:"id"`
	Language      string `json:"language"`
	Note          string `json:"note"`
	Status        string `json:"status"`
	BuildStdout   string `json:"build_stdout"`
	BuildStderr   string `json:"build_stderr"`
	BuildExitCode int    `json:"build_exit_code"`
	BuildTime     string `json:"build_time"`
	BuildMemory   int    `json:"build_memory"`
	BuildResult   string `json:"build_result"`
	Stdout        string `json:"stdout"`
	Stderr        string `json:"stderr"`
	ExitCode      int    `json:"exit_code"`
	Time          string `json:"time"`
	Memory        int    `json:"memory"`
	Connections   int    `json:"connections"`
	Result        string `json:"result"`
}

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

// MakeAttachments -
func (e *ExecutionResult) MakeAttachments() *[]*Attachment {
	var attachments []*Attachment

	buildMessage := message{status: status(e.BuildResult), time: e.BuildTime, memory: e.BuildMemory}
	buildAttachment := Attachment{Title: string("[BUILD " + strings.ToUpper(e.BuildResult) + "]")}
	if buildMessage.status == isSuccess {
		buildMessage.output = e.BuildStdout
		buildAttachment.Color = string(good)
		buildAttachment.Text = buildMessage.build()
		attachments = append(attachments, &buildAttachment)

		execMessage := message{status: status(e.Result), time: e.Time, memory: e.Memory}
		execAttachment := Attachment{Title: string("[EXEC " + strings.ToUpper(e.Result) + "]")}
		if execMessage.status == isSuccess {
			execMessage.output = e.Stdout
			execAttachment.Color = string(good)

		} else if execMessage.status == isFailure {
			execMessage.output = e.Stderr
			execAttachment.Color = string(warning)

		} else {
			execMessage.output = e.Stderr
			execAttachment.Color = string(danger)
		}

		execAttachment.Text = execMessage.build()
		attachments = append(attachments, &execAttachment)

	} else if buildMessage.status == isFailure {
		buildMessage.output = e.BuildStderr
		buildAttachment.Color = string(warning)
		buildAttachment.Text = buildMessage.build()
		attachments = append(attachments, &buildAttachment)

	} else {
		buildMessage.output = e.BuildStderr
		buildAttachment.Color = string(danger)
		buildAttachment.Text = buildMessage.build()
		attachments = append(attachments, &buildAttachment)
	}

	return &attachments
}

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
