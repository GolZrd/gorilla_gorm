package controller

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) handler {
	return handler{db}
}

// home handler
func (h handler) Home(w http.ResponseWriter, r *http.Request) {
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

func (h handler) GetNotes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Здесь будут все записи из базы данных"))
}

func (h handler) GetNote(w http.ResponseWriter, r *http.Request) {
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
