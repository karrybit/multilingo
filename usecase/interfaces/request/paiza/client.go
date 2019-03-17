package paiza

import (
	"github.com/TakumiKaribe/multilingo/entity/paiza"
)

type Client interface {
	Request(string, string) (*paiza.Result, error)
}
