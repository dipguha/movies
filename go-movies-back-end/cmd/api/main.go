package main

import (
	"backend/internal/repository"
	"backend/internal/repository/dbrepo"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const port = 8080

// *sql.DB - pointer to a poll of DB connection
// DB     *sql.DB change as part of 52
type application struct {
	DSN    string
	Domain string
	DB     repository.DataBaseRepo
}

func main() {

	//set application config
	var app application

	//read from command line
	//&app.DSN refers to the memory address and stores in p *string which is memory address of string variable
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	flag.Parse()

	log.Println("***** main - app.DSN *****: ", app.DSN)
	//connect to the db
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	log.Println("***** main - app.DB *****:", app.DB)
	defer app.DB.Connection().Close()

	app.Domain = "example.com"

	log.Println("***** main - Starting the application on port: ", port)

	//http.HandleFunc("/", Hello) - don't need any more

	//start a web server
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}

}
