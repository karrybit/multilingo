package interfaces

import (
	"io"
	"net/http"
	"time"

	"github.com/TakumiKaribe/multilingo/entity/paiza"
	"github.com/TakumiKaribe/multilingo/entity/slack"
	log "github.com/sirupsen/logrus"
)

type Request interface {
	// Clientと引数を受け取る
	Get(Method, string) (*io.ReadCloser, error)
}

type Method string

const (
	Get  Method = http.MethodGet
	Post Method = http.MethodPost
)

type Reqeuster struct {
	httpClient *http.Client
}

func NewRequester() *Reqeuster {
	return &Reqeuster{httpClient: &http.Client{Timeout: time.Duration(10) * time.Second}}
}

func (r *Reqeuster) Request(method Method, urlString string, body io.Reader, header map[string]string) (io.ReadCloser, error) {
	req, err := http.NewRequest(string(method), urlString, body)

	for k, v := range header {
		req.Header.Set(k, v)
	}

	log.Printf("⚡️  %s", req.URL)
	resp, err := r.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

type PaizaClient interface {
	Request(string, string) (*paiza.Result, error)
}

type SlackClient interface {
	Notify(*slack.RequestBody) error
}
