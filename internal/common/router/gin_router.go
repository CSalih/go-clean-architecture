package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type ginRouter struct {
	router *gin.Engine
}

func NewGinRouter() Router {
	return &ginRouter{
		router: gin.Default(),
	}
}

func (r *ginRouter) Get(relativePath string, handler HandlerFunc) {
	r.router.GET(ginPathAdapter(relativePath), r.ginHandleAdapter(handler))
}

func (r *ginRouter) Post(relativePath string, handler HandlerFunc) {
	r.router.POST(ginPathAdapter(relativePath), r.ginHandleAdapter(handler))
}

func (r *ginRouter) Put(relativePath string, handler HandlerFunc) {
	r.router.PUT(ginPathAdapter(relativePath), r.ginHandleAdapter(handler))
}

func (r *ginRouter) ginHandleAdapter(handler HandlerFunc) func(*gin.Context) {
	return func(ctx *gin.Context) {
		params := make(map[string]string)
		for i := 0; i < len(ctx.Params); i++ {
			params[ctx.Params[i].Key] = ctx.Params[i].Value
		}
		handler(&Context{
			Writer:  ctx.Writer,
			Request: ctx.Request,
			Params:  params,
		})
	}
}

func ginPathAdapter(relativePath string) string {
	path := strings.ReplaceAll(relativePath, "{", ":")
	path = strings.ReplaceAll(path, "}", "")
	return path
}

func (r *ginRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}
