// Package negronibugsnag provides a negroni middleware that reports exceptions
// to bugsnag.
package negronibugsnag // import "fknsrs.biz/p/negroni-bugsnag"

import (
	"net/http"

	"fknsrs.biz/p/bugsnag"
)

type Middleware struct {
	c *bugsnag.Client
}

// New returns a new instance of the middleware, given a *bugsnag.Client
func New(c *bugsnag.Client) *Middleware {
	return &Middleware{c: c}
}

func (m *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	defer m.c.ReportPanic()
	next(w, r)
}
