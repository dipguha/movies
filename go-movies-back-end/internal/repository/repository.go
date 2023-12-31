package repository

import (
	"backend/internal/models"
	"database/sql"
)

type DataBaseRepo interface {
	Connection() *sql.DB
	AllMovies() ([]*models.Movie, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(ID int) (*models.User, error)
}
