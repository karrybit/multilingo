package model

type Status struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

func (cr *Status) Log() {}
