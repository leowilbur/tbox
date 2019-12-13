package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/leowilbur/tbox/pkg/rest"
)

var (
	bind         = flag.String("bind", ":8080", "address which we should bind to")
	dbConnection = flag.String("dbConnection", "root:root@tcp(127.0.0.1:3306)/tbox?timeout=360s&multiStatements=true&parseTime=true", "address which database connection string")
)

func main() {
	log.Println("connecting to mysql database")
	db, err := sql.Open("mysql", *dbConnection)

	router, err := rest.New(
		db,
	)
	if err != nil {
		panic(err)
	}
	panic(http.ListenAndServe(*bind, router))
}
