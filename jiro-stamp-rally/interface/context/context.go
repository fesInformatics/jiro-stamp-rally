package context

import (
	"encoding/json"
	"net/http"
	"fmt"
	_context "context"
)

type Context struct {
	w http.ResponseWriter
	r  *http.Request
}

func (c Context) JSON (v interface{}) {
	if err := json.NewEncoder(c.w).Encode(v); err != nil {
		fmt.Printf("エンコードエラー")
	}
}

func (c Context) CreateContext () _context.Context {
	return c.r.Context()
}

func NewContext (w http.ResponseWriter, r  *http.Request) Context {
	return Context{
		w: w,
		r: r,
	}
}

