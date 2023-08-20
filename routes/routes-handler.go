package routes

import (
	"github.com/gorilla/mux"
	"github.com/zohaibsoomro/book-server-postgresq/controller"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/books", controller.GetAllBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", controller.GetBookById).Methods("GET")
	r.HandleFunc("/api/book", controller.CreateBook).Methods("POST")
	r.HandleFunc("/api/book/{id}", controller.UpdateBookById).Methods("PUT")
	r.HandleFunc("/api/book/{id}", controller.DeleteBookById).Methods("DELETE")
	return r
}
