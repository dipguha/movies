package main

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func openDB(dsn string) (*sql.DB, error) {
	log.Println("***** openDB dsn *****:", dsn)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	log.Println("***** openDB after sql Open *****: ", db)

	err = db.Ping()
	if err != nil {
		log.Println("***** Ping err *****: ", err)
		return nil, err
	}

	log.Println("***** openDB after Ping *****:", db)

	return db, nil
}

func (app *application) connectToDB() (*sql.DB, error) {
	log.Println("***** connectToDB start *****:", app)
	connection, err := openDB(app.DSN)
	if err != nil {
		return nil, err
	}

	log.Println("***** Connetced to Postgres *****:")
	return connection, nil
}
