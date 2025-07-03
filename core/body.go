package core

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

type Request struct {
	method   string
	path     string
	authType AuthType
	query    url.Values
	form     url.Values
	header   http.Header
	body     io.Reader
}

func (r *Request) Set(key string, value any) *Request {
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

func (r *Request) GetQuery(key string) string {
	return r.query.Get(key)
}

type response struct {
	status    int
	err       error
	rawBody   []byte
	rawHeader http.Header
}

type WsRequest struct {
	Id       string         `json:"id"`
	Method   string         `json:"method"`
	Params   map[string]any `json:"params,omitempty"`
	AuthType AuthType       `json:"-"`
}

func (r *WsRequest) Set(key string, value any) *WsRequest {
	r.Params[key] = value
	return r
}

func (r *WsRequest) Get(key string) any {
	if r.Params == nil {
		return nil
	}
	return r.Params[key]
}
