package main

import (
	"net/http"

	"github.com/LostLaser/TESTWEB/books"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/create", books.CreateController)
	r.HandleFunc("/books/create_handle", books.Create)
	r.HandleFunc("/books/search/{isbn}", books.Search)
	r.HandleFunc("/books", books.GetAll)
	r.HandleFunc("/books/delete", books.Delete)

	http.ListenAndServe(":8080", r)
}
