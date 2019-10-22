package main

import (
	"net/http"

	"github.com/LostLaser/TestWeb/controllers"
	"github.com/LostLaser/TestWeb/views"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books", views.Home)
	r.HandleFunc("/books/create_handle", controllers.CreateController)
	r.HandleFunc("/books/create", views.Create)
	r.HandleFunc("/books/delete_handle", controllers.DeleteController)

	http.ListenAndServe(":8080", r)
}
