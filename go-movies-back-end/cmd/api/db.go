package main

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func openDB(dsn string) (*sql.DB, error) {
	log.Println("***** db - openDB dsn *****:", dsn)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	log.Println("***** db - openDB after sql Open *****: ", db)

	err = db.Ping()
	if err != nil {
		log.Println("***** db - Ping err *****: ", err)
		return nil, err
	}

	log.Println("***** db - openDB after Ping *****:", db)

	return db, nil
}

func (app *application) connectToDB() (*sql.DB, error) {
	log.Println("***** db - connectToDB start *****:", app)
	connection, err := openDB(app.DSN)
	if err != nil {
		return nil, err
	}

	log.Println("***** db - Connetced to Postgres *****:")
	return connection, nil
}
