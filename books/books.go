package books

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

// Book contains information relevant to an individual book
type Book struct {
	Title string
	ISBN  string
	Genre string
}

// BookPageData is a response body for books
type BookPageData struct {
	PageTitle string
	Library   []Book
}

var library []Book

// CreateController renders book creation form
func CreateController(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("assets/layout.html", "assets/books/create.html"))
	tmpl.ExecuteTemplate(w, "layout", "")
}

// Create makes a book with the specified title and ISBN
func Create(w http.ResponseWriter, r *http.Request) {
	newbook := Book{}
	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// get the book
	newbook.Title = r.Form.Get("book_title")

	// navigate to the page
	newbook.ISBN = r.Form.Get("book_isbn")
	library = append(library, newbook)

	http.Redirect(w, r, "/books", http.StatusSeeOther)
}

// Search finds a book with the specified title
func Search(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var retVal Book
	isbn := vars["isbn"]

	for _, element := range library {
		if element.ISBN == isbn {
			retVal = element
		}
	}

	fmt.Fprintf(w, "Found %s with ISBN %s", retVal.Title, retVal.ISBN)
}

// GetAll retrieves all of the books in the library
func GetAll(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("assets/layout.html", "assets/index.html"))
	data := BookPageData{
		PageTitle: "Library",
		Library:   library,
	}

	tmpl.ExecuteTemplate(w, "layout", data)
}

// Delete removes a single book from the library
func Delete(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	isbn := r.Form.Get("isbn")
	deleteIndex := -1

	for index, element := range library {
		if element.ISBN == isbn {
			deleteIndex = index
		}
	}
	if deleteIndex >= 0 {
		library = append(library[:deleteIndex], library[deleteIndex+1:]...)
	}

	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
