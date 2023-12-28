package dbrepo

import (
	"backend/internal/models"
	"context"
	"database/sql"
	"log"
	"time"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

const dbTimeOut = time.Second * 3

func (m *PostgresDBRepo) Connection() *sql.DB {
	log.Println("***** dbrepo - Connection m - *****: ", m)
	return m.DB
}

func (m *PostgresDBRepo) AllMovies() ([]*models.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	log.Println("***** dbrepo - AllMovies m - *****: ", m)
	log.Println("***** dbrepo - AllMovies ctx - *****: ", ctx)
	defer cancel()

	query := `select id, title, release_date, runtime, mpaa_rating, description, coalesce(image,''),created_at, updated_at
			  from   movies
			  order by title`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*models.Movie

	for rows.Next() {
		var movie models.Movie
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.ReleaseDate,
			&movie.RunTime,
			&movie.MPAARating,
			&movie.Description,
			&movie.Image,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		log.Println("***** dbrepo - AllMovies movie - *****: ", movie)
		movies = append(movies, &movie)
	}
	log.Println("***** dbrepo - AllMovies movies - *****: ", movies)
	return movies, nil
}
