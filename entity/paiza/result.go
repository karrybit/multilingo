package paiza

// Result -
type Result struct {
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
