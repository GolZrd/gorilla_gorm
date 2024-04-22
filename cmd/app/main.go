package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"gorilla_gorm/internal/controller"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", controller.Home)
	router.HandleFunc("/notes", controller.GetNotes)
	router.HandleFunc("/note/{id:[0-9]+}", controller.GetNote).Methods("GET")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static/"))))

	logRouter := handlers.LoggingHandler(os.Stdout, router)

	//router.HandleFunc("/notes/create", controller.CreateNote).Methods("POST")
	//router.HandleFunc("/note/{id:[0-9]+}/update", controller.UpdateNote).Methods("POST")
	//router.HandleFunc("/note/{id:[0-9]+}/delete", controller.Delete).Methods("POST")

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", logRouter)
}
