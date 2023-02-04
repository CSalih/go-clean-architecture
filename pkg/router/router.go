package router

import "net/http"

type HandlerFunc func(ctx *Context)

// Router is a lightweight abstraction for a http server.
type Router interface {
	Get(relativePath string, handler HandlerFunc)
	Post(relativePath string, handler HandlerFunc)
	Put(relativePath string, handler HandlerFunc)

	// ServeHTTP implements the http.Handler interface.
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}
