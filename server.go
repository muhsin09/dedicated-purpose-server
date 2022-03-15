package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"temp-ledger/handlers"
)

type Args struct {
	port string
}

func Start(args Args) error {
	router := mux.NewRouter().PathPrefix("/api/v1/").Subrouter()

	handler := handlers.NewServerHandler()
	RegisterAllRoutes(router, handler)

	log.Println("Starting server at port : ", args.port)
	return http.ListenAndServe(args.port, router)

}

func RegisterAllRoutes(router *mux.Router, handler handlers.IServerHandler) {

	// set content type
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	router.HandleFunc("/addMessage", handler.AddMessage).Methods(http.MethodGet)
	router.HandleFunc("/pollMessage", handler.PollMessage).Methods(http.MethodGet)

}
