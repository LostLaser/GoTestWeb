package views

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/LostLaser/TestWeb/models"
)

// Create renders book creation form
func Create(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("assets/layout.html", "assets/books/create.html"))
	tmpl.ExecuteTemplate(w, "layout", "")
}

// Home retrieves all of the books in the library
func Home(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	tmpl := template.Must(template.ParseFiles("assets/layout.html", "assets/index.html"))
	returnBooks := []models.Book{}
	q := r.Form.Get("search_isbn")

	for _, element := range models.Library {
		if q == "" || element.ISBN == q || element.Title == q {
			returnBooks = append(returnBooks, element)
		}
	}

	data := models.BookPageData{
		PageTitle: "Library",
		Library:   returnBooks,
	}

	tmpl.ExecuteTemplate(w, "layout", data)
}
