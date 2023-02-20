package app

import (
	"fmt"
	"net/http"
)

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

func (router *Router) FindHandler(method string, path string) (http.HandlerFunc, bool, bool) {
	methods, exists := router.rules[path]
	handler, methodExists := methods[method]
	return handler, exists, methodExists
}

func (router *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, exists, methodExists := router.FindHandler(request.Method, request.URL.Path)
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s not Found", request.URL.Path)
		return
	}

	if !methodExists {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	handler(w, request)
}
