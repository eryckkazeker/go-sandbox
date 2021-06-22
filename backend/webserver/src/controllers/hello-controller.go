package controllers

import (
	"fmt"
	"go-sandbox/src/hello"
	"net/http"

	"github.com/gorilla/mux"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello go-sandbox")
}

func HelloFromInterface(w http.ResponseWriter, r *http.Request) {
	var writer hello.HelloWriter
	urlVars := mux.Vars(r)
	id := urlVars["helloId"]
	if id == "1" {
		writer = &hello.Hello1{}
	}
	if id == "2" {
		writer = &hello.Hello2{}
	}

	fmt.Fprint(w, writer.WriteHello())
}
