package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	connStr := "postgres://zohaib:zhs@1234@postgressql-crud-api.postgres.database.azure.com/booksdb?port=5432&sslmode=verify-full&connect_timeout=20"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func PingDB(s *sql.DB) {
	err := s.Ping()
	if err != nil {
		fmt.Println("Connection failed!")
		log.Fatal(err)
	}
	fmt.Println("Connection successful.")
}
