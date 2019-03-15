package entity

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

	if e.BuildResult == "" {
		attachments = append(attachments, e.makeExecResultAttachment())

	} else {
		attachments = *e.makeBuildResultAttachment()
	}

	return &attachments
}

func (e *ExecutionResult) makeExecResultAttachment() *Attachment {
	attachment := Attachment{Title: fmt.Sprintf("[EXEC %s]", strings.ToUpper(e.Result))}
	message := message{status: status(e.Result), time: e.Time, memory: e.Memory}

	if message.status == isSuccess {
		message.output = e.Stdout
		attachment.Color = string(good)

	} else if message.status == isFailure {
		message.output = e.Stderr
		attachment.Color = string(warning)

	} else {
		message.output = e.Stderr
		attachment.Color = string(danger)
	}

	attachment.Text = message.build()
	return &attachment
}

func (e *ExecutionResult) makeBuildResultAttachment() *[]*Attachment {
	var attachments []*Attachment

	attachment := Attachment{Title: fmt.Sprintf("[BUILD %s]", strings.ToUpper(e.BuildResult))}
	message := message{status: status(e.BuildResult), time: e.BuildTime, memory: e.BuildMemory}

	if message.status == isSuccess {
		message.output = e.BuildStdout
		attachment.Color = string(good)

	} else if message.status == isFailure {
		message.output = e.BuildStderr
		attachment.Color = string(warning)

	} else {
		message.output = e.BuildStderr
		attachment.Color = string(danger)
	}

	attachment.Text = message.build()
	attachments = append(attachments, &attachment)

	if message.status == isSuccess {
		attachments = append(attachments, e.makeExecResultAttachment())
	}

	return &attachments
}
