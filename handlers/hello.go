package handlers

import (
	"net/http"
	"log"
	"fmt"
	"io/ioutil"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello{
	return &Hello{}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello World")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "An error has occured", http.StatusBadRequest)
		return
	}

	log.Printf("Request Data: %s\n", d)

	fmt.Fprintf(w, "Hello %s", d)
}