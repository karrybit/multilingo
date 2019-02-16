package model

type DetailResponse struct {
	ID              string `json:"id"`
	Language        string `json:"language"`
	Note            string `json:"note"`
	Status          string `json:"status"`
	Build_stdout    string `json:"build_stdout"`
	Build_stderr    string `json:"build_stderr"`
	Build_exit_code int    `json:"build_exit_code"`
	Build_time      string `json:"build_time"`
	Build_memory    int    `json:"build_memory"`
	Build_result    string `json:"build_result"`
	Stdout          string `json:"stdout"`
	Stderr          string `json:"stderr"`
	Exit_code       int    `json:"exit_code"`
	Time            string `json:"time"`
	Memory          int    `json:"memory"`
	Connections     int    `json:"connections"`
	Result          string `json:"result"`
}
