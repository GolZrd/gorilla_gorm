package controller

import (
	"gorilla_gorm/internal/database"
	"html/template"
	"log"
	"net/http"
	"time"

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

	var Notes []database.Note
	h.DB.Order("Created desc").Find(&Notes)

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

	err = ts.Execute(w, Notes)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h handler) GetNotes(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"web/html/notesPage.html",
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

func (h handler) GetNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	nota := database.Note{}

	result := h.DB.Raw("SELECT * FROM notes WHERE id = ?", id).Scan(&nota)
	if result.RowsAffected == 0 || result.Error != nil { // Пока непонятно, потому что время отображается, попрубую так, если не получится,
		// то значит, надо переделать, например, если ID равно 0
		log.Println(result.Error)
		http.Error(w, http.StatusText(404), http.StatusNotFound)
	}

	files := []string{
		"web/html/showPage.html",
		"web/html/basePage.html",
		"web/html/footer.html",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, nota)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}

// if nota.ID == 0 {
// 	log.Println(err)
// 	http.Error(w, http.StatusText(404), http.StatusNotFound)
// }

// Создание заметки
func (h handler) CreateNote(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

		title := r.FormValue("title")
		content := r.FormValue("content")

		newNote := database.Note{Title: title, Content: content, Created: time.Now()}
		h.DB.Create(&newNote)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	} else {
		files := []string{
			"web/html/createPage.html",
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

}

func (h handler) UpdateNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	upNote := database.Note{}
	h.DB.First(&upNote, id)

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

		upNote.Title = r.FormValue("title")
		upNote.Content = r.FormValue("content")

		h.DB.Save(&upNote)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	} else {
		files := []string{
			"web/html/editPage.html",
			"web/html/basePage.html",
			"web/html/footer.html",
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		err = ts.Execute(w, upNote)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

func (h handler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	res := h.DB.Delete(&database.Note{}, id)
	if res.Error != nil {
		log.Println(res.Error)
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		return
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func (h handler) CreateTest(w http.ResponseWriter, r *http.Request) {

	newNote := database.Note{Title: "Test", Content: "Test", Created: time.Now()}
	h.DB.Create(&newNote)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
