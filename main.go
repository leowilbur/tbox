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
	dbConnection = flag.String("dbConnection", "", "address which database connection string")
)

func main() {
	log.Println("connecting to mysql database")
	db, err := sql.Open("mysql", "rdsuser:DEV7ywDRDwPKTsya@tcp(redisys-dev.cm6yeh8nkhqx.ap-southeast-2.rds.amazonaws.com:3306)/_test_migration?timeout=360s&multiStatements=true&parseTime=true")

	router, err := rest.New(
		db,
	)
	if err != nil {
		panic(err)
	}
	panic(http.ListenAndServe(":8080", router))
}
