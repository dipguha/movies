package repository

import (
	"backend/internal/models"
	"database/sql"
)

type DataBaseRepo interface {
	Connection() *sql.DB
	AllMovies() ([]*models.Movie, error)
}
