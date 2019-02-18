package model

// Status -
type Status struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

// Log is standard output of all properties
func (cr *Status) Log() {}
