package main

import (
	"log"
	"net/http"

	"github.com/zohaibsoomro/book-server-postgresq/config"
	"github.com/zohaibsoomro/book-server-postgresq/model"
	"github.com/zohaibsoomro/book-server-postgresq/routes"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()
	config.PingDB(db)
	model.SetDBClient(db)

	r := routes.RegisterRoutes()
	log.Fatal(http.ListenAndServe(":8080", r))

}
