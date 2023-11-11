package context

import (
	_context "context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Context struct {
	w http.ResponseWriter
	r  *http.Request
}

func (c Context) JSON (code int, v any) {
	c.w.Header().Set("Contetnt-Type","application/json; charset=utf-8")
	c.w.WriteHeader(code)
	if v != nil {
		if err := json.NewEncoder(c.w).Encode(v); err != nil {
			fmt.Printf("エンコードエラー")
		}
	}
}

func (c Context) CreateContext () _context.Context {
	return c.r.Context()
}

func (c Context) GetHttpMethod() string{
	return c.r.Method
}

func (c Context)GetBody() io.ReadCloser {
	return c.r.Body
}

func NewContext (w http.ResponseWriter, r  *http.Request) Context {
	return Context{
		w: w,
		r: r,
	}
}

