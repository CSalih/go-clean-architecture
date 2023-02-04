package router

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type chiRouter struct {
	router *chi.Mux
}

func NewChiRouter() Router {
	return &chiRouter{
		router: chi.NewRouter(),
	}
}

func (r *chiRouter) Get(pattern string, handler HandlerFunc) {
	r.router.Get(pattern, r.chiHandleAdapter(handler))
}

func (r *chiRouter) Post(pattern string, handler HandlerFunc) {
	r.router.Post(pattern, r.chiHandleAdapter(handler))
}

func (r *chiRouter) Put(pattern string, handler HandlerFunc) {
	r.router.Put(pattern, r.chiHandleAdapter(handler))
}

func (r *chiRouter) chiHandleAdapter(handler HandlerFunc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		urlParams := chi.RouteContext(r.Context()).URLParams
		params := make(map[string]string)
		for i := 0; i < len(urlParams.Keys); i++ {
			params[urlParams.Keys[i]] = urlParams.Values[i]
		}
		handler(&Context{
			Writer:  w,
			Request: r,
			Params:  params,
		})
	}
}

func (r *chiRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}
