package main

import (
	"net/http"

	"github.com/LostLaser/TestWeb/controllers"
	"github.com/LostLaser/TestWeb/views"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books", views.HomeBook)
	r.HandleFunc("/books/create_handle", controllers.CreateBookController)
	r.HandleFunc("/books/create", views.CreateBook)
	r.HandleFunc("/books/delete_handle", controllers.DeleteBookController)

	http.ListenAndServe(":8080", r)
}
