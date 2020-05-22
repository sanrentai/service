package main

import (
	"fmt"
	"net/http"

	"github.com/sanrentai/service"
)

func main() {
	prg := &service.Program{
		Name:        "test",
		DisplayName: "TEST",
		Description: "this is a test",
		Runfunc:     Run,
		Stopfunc:    Close,
	}
	service.Run(prg)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	service.Info("hello world")
	fmt.Fprintf(w, "hello world")
}

var server *http.Server

func Run() error {
	http.HandleFunc("/", indexHandler)
	server = &http.Server{Addr: ":8000"}
	return server.ListenAndServe()
}

func Close() error {

	service.Info("Close")
	server.Close()
	return nil
}
