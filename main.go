package main

import (
	"log"
	"net/http"

	"github.com/jscastaneda-esp/server-vanilla-go/app"
)

const (
	maxWorkers   = 4
	maxQueueSize = 20
	port         = ":8081"
)

func main() {
	server := app.NewServer(port)
	server.Handle(http.MethodGet, "/", app.HandleRoot)
	server.Handle(http.MethodGet, "/api", app.HandleHome)
	server.Handle(http.MethodPost, "/api", server.AddMiddleware(app.HandleHome, app.CheckAuth(), app.Logging()))
	server.Handle(http.MethodPost, "/create", app.UserPostRequest)

	jobQueue := make(chan app.Job, maxQueueSize)
	dispatcher := app.NewDispatcher(jobQueue, maxWorkers)
	dispatcher.Run()
	server.Handle(http.MethodPost, "/fib", func(w http.ResponseWriter, r *http.Request) {
		app.RequestHandler(w, r, jobQueue)
	})

	log.Fatal(server.Listen())
}
