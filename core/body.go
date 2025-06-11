package core

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

type request struct {
	method   string
	path     string
	authType AuthType
	query    url.Values
	form     url.Values
	header   http.Header
	body     io.Reader
}

type response struct {
	status    int
	err       error
	rawBody   []byte
	rawHeader http.Header
}

func (r *request) set(key string, value any) *request {
	if r.query == nil {
		r.query = url.Values{}
	}
	switch reflect.TypeOf(value).Kind() {
	case reflect.Slice:
		if elems, ok := value.([]string); ok {
			r.query.Set(key, `["`+strings.Join(elems, `","`)+`"]`)
		} else {
			r.query.Set(key, fmt.Sprintf("%v", value))
		}
	default:
		r.query.Set(key, fmt.Sprintf("%v", value))
	}
	return r
}

type wsRequest struct {
	Id       string         `json:"id"`
	Method   string         `json:"method"`
	Params   map[string]any `json:"params,omitempty"`
	AuthType AuthType       `json:"-"`
}
