package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zohaibsoomro/book-server-postgresq/model"
)

var CreateBook = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	b := model.Book{}
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("Request not have valid body.\n"))
		w.Write([]byte(err.Error()))
		return
	}

	if _, err := b.CreateBookInDB(); err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("Book creation failed.\n"))
		w.Write([]byte(err.Error()))
		return
	}
	s, _ := json.Marshal(b)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Book created\n"))
	w.Write(s)
}

var GetAllBooks = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	books, err := model.GetAllBooksFromDB()
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	s, err := json.Marshal(books)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(s)
}

var GetBookById = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("Book id is invalid.\n"))
		w.Write([]byte(err.Error()))
		return
	}
	b, err := model.GetBookByIdFromDb(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Book id not found!\n"))
		w.Write([]byte(err.Error()))
		return
	}
	s, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(s)
}

var UpdateBookById = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("Book id is invalid.\n"))
		w.Write([]byte(err.Error()))
		return
	}
	b, err := model.GetBookByIdFromDb(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Id not found!\n"))
		w.Write([]byte(err.Error()))
		return
	}
	book := model.Book{}
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("Request not have valid body.\n"))
		w.Write([]byte(err.Error()))
		return
	}
	if b.Author != book.Author {
		b.Author = book.Author
	}
	if b.Title != book.Title {
		b.Title = book.Title
	}
	if b.PublishDate != book.PublishDate {
		b.PublishDate = book.PublishDate
	}
	if err = b.UpdateBookByIdFromDb(); err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("Update failed!\n"))
		w.Write([]byte(err.Error()))
		return
	}
	s, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Book Updated\n"))
	w.Write(s)
}
var DeleteBookById = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("Book id not formatted.\n"))
		w.Write([]byte(err.Error()))
		return
	}
	b, err := model.GetBookByIdFromDb(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Id not found!\n"))
		return
	}

	if err = b.DeleteBookByIdFromDb(); err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("Deletion failed!\n"))
		w.Write([]byte(err.Error()))
		return
	}
	s, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted book:\n"))
	w.Write(s)
}
