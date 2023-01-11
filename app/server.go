package app

import (
	"net/http"
)

type Server struct {
	port   string
	router *Router
}

func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

func (server *Server) Listen() error {
	http.Handle("/", server.router)
	err := http.ListenAndServe(server.port, nil)
	if err != nil {
		return err
	}

	return nil
}

func (server *Server) Handle(method string, path string, handler http.HandlerFunc) {
	if _, ok := server.router.rules[path]; !ok {
		server.router.rules[path] = make(map[string]http.HandlerFunc)
	}

	server.router.rules[path][method] = handler
}

func (server *Server) AddMiddleware(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}

	return handler
}
