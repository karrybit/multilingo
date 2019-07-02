package paiza

import (
	"multilingo/entity/paiza"
)

// Client -
type Client interface {
	Request(string, string, string) (*paiza.Result, error)
}
