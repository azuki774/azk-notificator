package server

import (
	"azk-notificator/internal/model"
	"azk-notificator/internal/telemetry"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

var ServerForHandler *Server // Server pointer for Handler
var decoder = schema.NewDecoder()

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK\n")
}

func enqueueHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	qk := vars["queue_kind"]

	var q model.Queue
	enqHeader := &model.EnqueueHeader{}
	if err := decoder.Decode(enqHeader, r.URL.Query()); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "internal error koko: %v\n", err)
		return
	}

	q.From = enqHeader.From
	q.Title = enqHeader.Title
	q.To = enqHeader.To

	switch qk {
	case model.QueueKindOnlyLogStr:
		q.Kind = model.QueueKindOnlyLog
	case model.QueueKindEmailStr:
		q.Kind = model.QueueKindEmail
	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "queue kind is not set.\n")
		return
	}

	// Get body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "internal error: %v\n", err)
		return
	}
	defer r.Body.Close()
	q.Body = body

	ctx := telemetry.NewCtxWithSpanID()
	err = ServerForHandler.Enqueue(ctx, q)

	if err != nil {
		if errors.Is(err, model.ErrOverCapacity) {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintf(w, "queue over capacity: %v\n", err)
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "internal error: %v\n", err)
	}

	w.WriteHeader(http.StatusNoContent)
}

func dequeueHandler(w http.ResponseWriter, r *http.Request) {
	ctx := telemetry.NewCtxWithSpanID()
	q, err := ServerForHandler.Dequeue(ctx)
	if err != nil {
		if errors.Is(err, model.ErrQueueNotFound) {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "internal error: %v\n", err)
		return
	}

	b, err := json.Marshal(q)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "internal error: %v\n", err)
		return
	}

	fmt.Fprintln(w, string(b))
	w.WriteHeader(http.StatusOK)
}

func NewHandler() (r *mux.Router) {
	r = mux.NewRouter()

	// Add Hundler
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/enqueue/{queue_kind}", enqueueHandler).Methods("POST")
	r.HandleFunc("/dequeue", dequeueHandler).Methods("POST")
	return r
}
