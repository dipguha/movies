package main

import (
	"backend/internal/models"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello World from %s", app.Domain)
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "Active",
		Message: "Go movies up and running",
		Version: "1.0.0",
	}

	out, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	var movies []models.Movie

	rd, _ := time.Parse("2006-01-02", "2001-01-01")

	movie1 := models.Movie{
		ID:          1,
		Title:       "Movie 1",
		ReleaseDate: rd,
		RunTime:     111,
		MPAARating:  "R1",
		Description: "Description 1",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, movie1)

	rd, _ = time.Parse("2006-01-02", "2002-02-02")

	movie2 := models.Movie{
		ID:          2,
		Title:       "Movie 2",
		ReleaseDate: rd,
		RunTime:     222,
		MPAARating:  "R2",
		Description: "Description 2",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, movie2)

	out, err := json.Marshal(movies)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
