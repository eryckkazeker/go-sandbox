package memdb

import (
	"go-sandbox/src/models"
	"log"

	"github.com/google/uuid"
)

type MemoryDatabaseMovieDaoImpl struct {
	movies map[string]models.Movie
}

var movieDaoInstance *MemoryDatabaseMovieDaoImpl

func (m MemoryDatabaseMovieDaoImpl) GetInstance() *MemoryDatabaseMovieDaoImpl {
	if movieDaoInstance != nil {
		return movieDaoInstance
	}

	movieDaoInstance = &MemoryDatabaseMovieDaoImpl{movies: make(map[string]models.Movie)}
	return movieDaoInstance
}

func (m MemoryDatabaseMovieDaoImpl) AddMovie(movie models.Movie) (string, error) {
	id := uuid.New().String()
	log.Println(id)
	movie.Id = id
	m.movies[id] = movie
	return id, nil
}

func (m MemoryDatabaseMovieDaoImpl) ListMovies() ([]models.Movie, error) {
	list := make([]models.Movie, 0)
	for _, value := range m.movies {
		list = append(list, value)
	}
	return list, nil
}

func (m MemoryDatabaseMovieDaoImpl) GetMovieById(id string) (models.Movie, error) {
	return m.movies[id], nil
}

func (m MemoryDatabaseMovieDaoImpl) UpdateMovie(movie, newMovie models.Movie) (models.Movie, error) {
	movie.Title = newMovie.Title
	movie.Year = newMovie.Year
	m.movies[movie.Id] = movie
	return movie, nil
}

func (m MemoryDatabaseMovieDaoImpl) DeleteMovie(id string) error {
	delete(m.movies, id)
	return nil
}
