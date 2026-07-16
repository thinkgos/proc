package enum_spec

import (
	"cmp"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Loader struct {
	client *http.Client
}

func NewLoader() *Loader {
	return &Loader{}
}

func (l *Loader) SetHTTPClient(cl *http.Client) *Loader {
	l.client = cl
	return l
}

func (l *Loader) LoadFromData(data []byte) (*T, error) {
	var t T
	err := json.Unmarshal(data, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
func (l *Loader) LoadFromReader(r io.Reader) (*T, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return l.LoadFromData(data)
}

func (l *Loader) LoadFromFile(filename string) (*T, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return l.LoadFromData(data)
}

func (l *Loader) LoadFromURL(url string) (*T, error) {
	data, err := ReadFromHTTP(l.client, url)
	if err != nil {
		return nil, err
	}
	return l.LoadFromData(data)
}

func ReadFromHTTP(cl *http.Client, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
	if err != nil {
		return nil, err
	}
	cl = cmp.Or(cl, http.DefaultClient)
	resp, err := cl.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() //nolint:errcheck
	if resp.StatusCode > 399 {
		return nil, fmt.Errorf("error loading %q: request returned status code %d", url, resp.StatusCode)
	}
	return io.ReadAll(resp.Body)
}
