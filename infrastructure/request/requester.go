package request

import (
	"io"
	"log"
	"net/http"
	"time"

	interfacesRequest "github.com/TakumiKaribe/multilingo/usecase/interfaces/request"
)

type Reqeuster struct {
	httpClient *http.Client
}

func NewRequester() *Reqeuster {
	return &Reqeuster{httpClient: &http.Client{Timeout: time.Duration(10) * time.Second}}
}

func (r *Reqeuster) Request(method interfacesRequest.Method, urlString string, body io.Reader, header map[string]string) (io.ReadCloser, error) {
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
