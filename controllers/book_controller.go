package controllers

import (
	"fmt"
	"net/http"

	"github.com/LostLaser/TestWeb/models"
)

// DeleteBookController removes a single book from the library
func DeleteBookController(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	isbn := r.Form.Get("isbn")

	models.DeleteBook(isbn)

	http.Redirect(w, r, "/books", http.StatusSeeOther)
}

// CreateBookController makes a book with the specified title and ISBN
func CreateBookController(w http.ResponseWriter, r *http.Request) {
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

	err = models.AddBook(newbook)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
	}

	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
