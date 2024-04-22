package controller

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// home handler
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"web/html/homePage.html",
		"web/html/basePage.html",
		"web/html/footer.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func GetNotes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Здесь будут все записи из базы данных"))
}

func GetNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	w.Write([]byte("Запись с ID " + id))
}

//Создание заметки
// func CreateNote(w http.ResponseWriter, r *http.Request) {}

//Редактирование заметки
// func UpdateNote(w http.ResponseWriter, r *http.Request) {}

//Удаление заметки
// func Delete(w http.ResponseWriter, r *http.Request) {}
