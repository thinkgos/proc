package lookup

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestExtractor(t *testing.T) {
	var extractorTestTokenValue = "testTokenValue"

	var tests = []struct {
		name      string
		extractor Extractor
		headers   map[string]string
		query     url.Values
		cookie    map[string]string
		value     string
		err       error
	}{
		{
			name:      "header hit",
			extractor: HeaderExtractor{"token", ""},
			headers:   map[string]string{"token": extractorTestTokenValue},
			query:     nil,
			cookie:    nil,
			value:     extractorTestTokenValue,
			err:       nil,
		},
		{
			name:      "header miss",
			extractor: HeaderExtractor{"This-Header-Is-Not-Set", ""},
			headers:   map[string]string{"token": extractorTestTokenValue},
			query:     nil,
			cookie:    nil,
			value:     "",
			err:       ErrMissingValue,
		},

		{
			name:      "header filter",
			extractor: HeaderExtractor{"Authorization", "Bearer"},
			headers:   map[string]string{"Authorization": "Bearer " + extractorTestTokenValue},
			query:     nil,
			cookie:    nil,
			value:     extractorTestTokenValue,
			err:       nil,
		},
		{
			name:      "header filter miss",
			extractor: HeaderExtractor{"Authorization", "Bearer"},
			headers:   map[string]string{"Authorization": "Bearer   "},
			query:     nil,
			cookie:    nil,
			value:     "",
			err:       ErrMissingValue,
		},
		{
			name:      "argument hit",
			extractor: ArgumentExtractor("token"),
			headers:   map[string]string{},
			query:     url.Values{"token": {extractorTestTokenValue}},
			cookie:    nil,
			value:     extractorTestTokenValue,
			err:       nil,
		},
		{
			name:      "argument miss",
			extractor: ArgumentExtractor("token"),
			headers:   map[string]string{},
			query:     nil,
			cookie:    nil,
			value:     "",
			err:       ErrMissingValue,
		},
		{
			name:      "cookie hit",
			extractor: CookieExtractor("token"),
			headers:   map[string]string{},
			query:     nil,
			cookie:    map[string]string{"token": extractorTestTokenValue},
			value:     extractorTestTokenValue,
			err:       nil,
		},
		{
			name:      "cookie miss",
			extractor: ArgumentExtractor("token"),
			headers:   map[string]string{},
			query:     nil,
			cookie:    map[string]string{},
			value:     "",
			err:       ErrMissingValue,
		},
		{
			name:      "cookie miss",
			extractor: ArgumentExtractor("token"),
			headers:   map[string]string{},
			query:     nil,
			cookie:    map[string]string{"token": " "},
			value:     "",
			err:       ErrMissingValue,
		},
	}
	// Bearer token request
	for _, e := range tests {
		// Make request from test struct
		r := makeTestRequest("GET", "/", e.headers, e.cookie, e.query)

		// Test extractor
		token, err := e.extractor.ExtractValue(r)
		if token != e.value {
			t.Errorf("[%v] Expected token '%v'.  Got '%v'", e.name, e.value, token)
			continue
		}
		if err != e.err {
			t.Errorf("[%v] Expected error '%v'.  Got '%v'", e.name, e.err, err)
			continue
		}
	}
}

func makeTestRequest(method, path string, headers, cookie map[string]string, urlArgs url.Values) *http.Request { //nolint:unparam
	r, err := http.NewRequestWithContext(context.TODO(), method, fmt.Sprintf("%v?%v", path, urlArgs.Encode()), nil) //nolint:gocritic
	if err != nil {
		panic(err)
	}
	for k, v := range headers {
		r.Header.Set(k, v)
	}
	for k, v := range cookie {
		r.AddCookie(&http.Cookie{Name: k, Value: v})
	}
	return r
}
