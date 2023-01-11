package main

import (
	"log"
	"net/http"

	"github.com/jscastaneda-esp/server-vanilla-go/app"
)

func main() {
	server := app.NewServer(":3000")
	server.Handle(http.MethodGet, "/", app.HandleRoot)
	server.Handle(http.MethodGet, "/api", app.HandleHome)
	server.Handle(http.MethodPost, "/api", server.AddMiddleware(app.HandleHome, app.CheckAuth(), app.Logging()))
	server.Handle(http.MethodPost, "/create", app.UserPostRequest)
	log.Fatalln(server.Listen())
}
