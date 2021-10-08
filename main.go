package main

import (
	"net/http"

	"internal/taskstore"
)

type taskServer struct {
	store *taskstore.TaskStore
}

func main() {
	mux := http.NewServeMux()
	server := NewTaskServer()
	mux.HandleFunc("/task/", server.TaskHandler)
}
