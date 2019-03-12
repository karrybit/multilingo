package model

import (
	"fmt"
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

// MakeAttachments -
func (e *ExecutionResult) MakeAttachments() *[]*Attachment {
	var attachments []*Attachment
	if e.BuildResult != "success" {
		buildFailedText := fmt.Sprintf("status: failed\ntime: %ssec.\nmemory used: %d bytes", e.BuildTime, e.BuildMemory)
		if len(e.BuildStderr) > 0 {
			buildFailedText += fmt.Sprintf("\nlog:\n```%s```", e.BuildStderr)
		}
		buildFailedAttachment := Attachment{Color: "warning", Title: "Build", Text: buildFailedText}
		attachments = append(attachments, &buildFailedAttachment)

	} else {
		buildSucceededText := fmt.Sprintf("status: succeeded\ntime: %ssec\nmemory used: %d bytes", e.BuildTime, e.BuildMemory)
		if len(e.BuildStdout) > 0 {
			buildSucceededText += fmt.Sprintf("\nlog:\n```%s```", e.BuildStdout)
		}
		buildSucceededAttachment := Attachment{Color: "good", Title: "Build", Text: buildSucceededText}
		attachments = append(attachments, &buildSucceededAttachment)

		if e.Result != "success" {
			execFailedText := fmt.Sprintf("status: failed\ntime: %ssec\nmemory used: %d bytes", e.Time, e.Memory)
			if len(e.Stderr) > 0 {
				execFailedText += fmt.Sprintf("\nlog:\n```%s```", e.Stderr)
			}
			attachment := Attachment{Color: "danger", Title: "Exec", Text: execFailedText}
			attachments = append(attachments, &attachment)

		} else {
			execSucceededText := fmt.Sprintf("status: succeeded\ntime: %ssec\nmemory used: %d bytes", e.Time, e.Memory)
			if len(e.Stdout) > 0 {
				execSucceededText += fmt.Sprintf("\nlog:\n```%s```", e.Stdout)
			}
			attachment := Attachment{Color: "good", Title: "Exec", Text: execSucceededText}
			attachments = append(attachments, &attachment)
		}
	}

	return &attachments
}
