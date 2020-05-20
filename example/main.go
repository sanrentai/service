package main

import (
	"fmt"
	"net/http"

	"github.com/sanrentai/service"
	"github.com/sanrentai/service/log"
)

func main() {
	prg := &service.Program{
		Name:        "test",
		DisplayName: "TEST",
		Description: "this is a test",
		Runfunc:     Run,
	}
	service.Run(prg)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("hello world")
	fmt.Fprintf(w, "hello world")
}

func Run() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}
