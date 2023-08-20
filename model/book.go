package model

import (
	"database/sql"
	"fmt"
	"strings"
)

type Book struct {
	ID          int64  `json:"id"`
	Title       string `json:"name"`
	Author      string `json:"author"`
	PublishDate string `json:"publish_date"`
}

var client *sql.DB

func SetDBClient(d *sql.DB) {
	client = d
}

func GetAllBooksFromDB() ([]Book, error) {
	queryStr := "Select * from books"
	res, err := client.Query(queryStr)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	var books []Book

	for res.Next() {
		b := Book{}
		res.Scan(&b.ID, &b.Title, &b.Author, &b.PublishDate)
		books = append(books, b)
	}
	return books, nil
}

func (b *Book) CreateBookInDB() (int64, error) {
	insertQuery := "INSERT INTO books(title,author,publish_date) Values($1,$2,$3) RETURNING id"
	err := client.QueryRow(insertQuery, b.Title, b.Author, b.PublishDate).Scan(&b.ID)
	if err != nil {
		return 0, err
	}
	fmt.Println("Book created with id", b.ID)
	return b.ID, nil
}

func GetBookByIdFromDb(id int64) (*Book, error) {
	b := Book{}
	query := "Select * from books Where id=$1"
	err := client.QueryRow(query, id).Scan(&b.ID, &b.Title, &b.Author, &b.PublishDate)
	if err != nil {
		if strings.EqualFold(err.Error(), "sql: no rows in result set") {
			return nil, fmt.Errorf("id %v not found", id)
		}
		return nil, err
	}
	return &b, nil
}

func (b *Book) UpdateBookByIdFromDb() error {
	updateQuery := "Update books set title=$2,author=$3,publish_date=$4 Where id=$1"
	res, err := client.Exec(updateQuery, b.ID, b.Title, b.Author, b.PublishDate)
	if err != nil {
		if strings.EqualFold(err.Error(), "sql: no rows in result set") {
			return fmt.Errorf("id %v not found", b.ID)
		}
		return err
	}
	count, _ := res.RowsAffected()
	fmt.Println("Rows updated:", count)
	return nil
}

func (b *Book) DeleteBookByIdFromDb() error {
	deleteQuery := "Delete from books Where id=$1"
	res, err := client.Exec(deleteQuery, b.ID)
	if err != nil {
		if strings.EqualFold(err.Error(), "sql: no rows in result set") {
			return fmt.Errorf("id %v not found", b.ID)
		}
		return err
	}
	count, _ := res.RowsAffected()
	fmt.Println("Rows deleted:", count)
	return nil
}
