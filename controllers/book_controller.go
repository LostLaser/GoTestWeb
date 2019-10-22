package controllers

import (
	"fmt"
	"net/http"

	"github.com/LostLaser/TestWeb/models"
)

// DeleteController removes a single book from the library
func DeleteController(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	isbn := r.Form.Get("isbn")
	deleteIndex := -1

	for index, element := range models.Library {
		if element.ISBN == isbn {
			deleteIndex = index
		}
	}
	if deleteIndex >= 0 {
		models.Library = append(models.Library[:deleteIndex], models.Library[deleteIndex+1:]...)
	}

	http.Redirect(w, r, "/books", http.StatusSeeOther)
}

// CreateController makes a book with the specified title and ISBN
func CreateController(w http.ResponseWriter, r *http.Request) {
	newbook := models.Book{}
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
	models.Library = append(models.Library, newbook)

	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
