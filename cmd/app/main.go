package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"gorilla_gorm/internal/controller"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", controller.Home)
	router.HandleFunc("/notes", controller.GetNotes)
	router.HandleFunc("/note/{id:[0-9]+}", controller.GetNote).Methods("GET")

	//router.HandleFunc("/notes/create", controller.CreateNote).Methods("POST")
	//router.HandleFunc("/note/{id:[0-9]+}/update", controller.UpdateNote).Methods("POST")
	//router.HandleFunc("/note/{id:[0-9]+}/delete", controller.Delete).Methods("POST")

	http.ListenAndServe(":8080", router)
}
