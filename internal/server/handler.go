package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK\n")
}

func enqueueHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "POST OK\n")
}

func NewHandler() (r *mux.Router) {
	r = mux.NewRouter()

	// Add Hundler
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/enqueue/", enqueueHandler).Methods("POST")
	return r
}
