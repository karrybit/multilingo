package interfaces

import (
	"io"
	"net/http"
)

type Method string

const (
	Get  Method = http.MethodGet
	Post Method = http.MethodPost
)

type Request interface {
	// Clientと引数を受け取る
	Get(Method, string) (*io.ReadCloser, error)
}
