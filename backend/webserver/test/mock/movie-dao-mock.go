package mock

import "go-sandbox/src/models"

type MovieDaoMock struct {
	AddMovieFunc     func(movie models.Movie) (string, error)
	ListMoviesFunc   func() ([]models.Movie, error)
	GetMovieByIdFunc func(id string) (models.Movie, error)
	UpdateMovieFunc  func(movie, newMovie models.Movie) (models.Movie, error)
	DeleteMovieFunc  func(id string) error
}

func (mock *MovieDaoMock) AddMovie(movie models.Movie) (string, error) {
	return mock.AddMovieFunc(movie)
}
func (mock *MovieDaoMock) ListMovies() ([]models.Movie, error) {
	return mock.ListMoviesFunc()
}
func (mock *MovieDaoMock) GetMovieById(id string) (models.Movie, error) {
	return mock.GetMovieByIdFunc(id)
}
func (mock *MovieDaoMock) UpdateMovie(movie, newMovie models.Movie) (models.Movie, error) {
	return mock.UpdateMovieFunc(movie, newMovie)
}
func (mock *MovieDaoMock) DeleteMovie(id string) error {
	return mock.DeleteMovieFunc(id)
}
