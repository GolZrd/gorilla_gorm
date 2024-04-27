package main

import (
	"fmt"
	"gorilla_gorm/internal/controller"
	"gorilla_gorm/internal/database"
	"os"

	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	DB := database.Init()
	h := controller.New(DB)

	router := mux.NewRouter()
	router.HandleFunc("/", h.Home)
	router.HandleFunc("/notes", h.GetNotes)
	router.HandleFunc("/note/{id:[0-9]+}", h.GetNote).Methods("GET")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static/"))))

	logRouter := handlers.LoggingHandler(os.Stdout, router)

	srv := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      logRouter,
	}

	//router.HandleFunc("/notes/create", controller.CreateNote).Methods("POST")
	//router.HandleFunc("/note/{id:[0-9]+}/update", controller.UpdateNote).Methods("POST")
	//router.HandleFunc("/note/{id:[0-9]+}/delete", controller.Delete).Methods("POST")

	fmt.Println("Server started on port 8080")
	srv.ListenAndServe()
}
