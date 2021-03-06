package views

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/LostLaser/TestWeb/models"
)

// CreateBook renders book creation form
func CreateBook(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("assets/layout.html", "assets/books/create.html"))
	tmpl.ExecuteTemplate(w, "layout", "")
}

// HomeBook retrieves all of the books in the library
func HomeBook(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	tmpl := template.Must(template.ParseFiles("assets/layout.html", "assets/index.html"))
	q := r.Form.Get("q")

	returnBooks := models.SearchLibrary(q, q)

	data := models.BookPageData{
		PageTitle: "Library",
		Library:   returnBooks,
	}

	tmpl.ExecuteTemplate(w, "layout", data)
}
