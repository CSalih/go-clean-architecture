package router

import "net/http"

type HandlerFunc func(ctx *Context)

// Router is a lightweight abstraction for a http server.
// It provides a simple way to register handlers for different
// HTTP endpoints and to serve them.
type Router interface {
	// Get registers a handler for the given relative path.
	Get(relativePath string, handler HandlerFunc)
	// Post registers a handler for the given relative path.
	Post(relativePath string, handler HandlerFunc)
	// Put registers a handler for the given relative path.
	Put(relativePath string, handler HandlerFunc)

	// ServeHTTP implements the http.Handler interface.
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}
