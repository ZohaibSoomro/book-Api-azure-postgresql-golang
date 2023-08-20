package routes

import (
	"github.com/gorilla/mux"
	"github.com/zohaibsoomro/book-server-postgresq/controller"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/books", controller.GetAllBooks).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/books/{id}", controller.GetBookById).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/book", controller.CreateBook).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/book/{id}", controller.UpdateBookById).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/book/{id}", controller.DeleteBookById).Methods("DELETE")
	return r
}
