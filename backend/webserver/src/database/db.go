package database

import (
	"go-sandbox/src/models"
)

type MovieDao interface {
	AddMovie(movie models.Movie) (string, error)
	ListMovies() ([]models.Movie, error)
	GetMovieById(id string) (models.Movie, error)
	UpdateMovie(movie, newMovie models.Movie) (models.Movie, error)
	DeleteMovie(id string) error
}

type UserDao interface {
	AddUser(user models.Login)
}
