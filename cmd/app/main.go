package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/notes", getNotes).Methods("GET")
	router.HandleFunc("/note/{id:[0-9]+}", getNote).Methods("GET")

	http.ListenAndServe(":8080", router)
}
