package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"go-sandbox/src/controllers"
	"go-sandbox/src/models"
	"go-sandbox/test/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMovieController_PostMovie_ShouldReturnMovieId_WhenSuccess(t *testing.T) {
	movie := &models.Movie{Title: "Anytitle", Year: 1999}
	movieJson, _ := json.Marshal(movie)
	reqBody := bytes.NewBuffer(movieJson)
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/movies", reqBody)

	var controller = &controllers.MovieController{
		DB: &mock.MovieDaoMock{
			AddMovieFunc: func(movie models.Movie) (string, error) {
				return "abc", nil
			},
		}}

	http.HandlerFunc(controller.PostMovie).ServeHTTP(rec, req)

	expectedBody := "abc"
	expectedStatusCode := http.StatusOK
	if expectedBody != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedBody, rec.Body.String())
	}
	if expectedStatusCode != rec.Code {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedStatusCode, rec.Code)
	}

}

func TestMovieController_PostMovie_ShouldReturnError_WhenErrorHappens(t *testing.T) {
	movie := &models.Movie{Title: "Anytitle", Year: 1999}
	movieJson, _ := json.Marshal(movie)
	reqBody := bytes.NewBuffer(movieJson)
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/movies", reqBody)

	var controller = &controllers.MovieController{
		DB: &mock.MovieDaoMock{
			AddMovieFunc: func(movie models.Movie) (string, error) {
				return "", errors.New("Error")
			},
		}}

	http.HandlerFunc(controller.PostMovie).ServeHTTP(rec, req)

	expectedBody := "Error saving Movie"
	expectedStatusCode := http.StatusInternalServerError
	if expectedBody != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedBody, rec.Body.String())
	}
	if expectedStatusCode != rec.Code {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedStatusCode, rec.Code)
	}
}

func TestMovieController_ListMovies_ShouldReturnArrayOfMovies(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/movies", nil)

	var controller = &controllers.MovieController{
		DB: &mock.MovieDaoMock{
			ListMoviesFunc: func() ([]models.Movie, error) {
				m1 := &models.Movie{Id: "123", Title: "Title1", Year: 1999}
				m2 := &models.Movie{Id: "456", Title: "Title2", Year: 2000}
				array := make([]models.Movie, 2)
				array[0] = *m1
				array[1] = *m2
				return array, nil
			},
		}}

	http.HandlerFunc(controller.ListMovies).ServeHTTP(rec, req)

	var resultArray []models.Movie

	json.Unmarshal(rec.Body.Bytes(), &resultArray)

	expectedArrayLength := 2
	expectedStatusCode := http.StatusOK
	if expectedArrayLength != len(resultArray) {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedArrayLength, len(resultArray))
	}
	if expectedStatusCode != rec.Code {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedStatusCode, rec.Code)
	}
}

func TestMovieController_ListMovies_ShouldReturnNoContent_WhenNoContent(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/movies", nil)

	var controller = &controllers.MovieController{
		DB: &mock.MovieDaoMock{
			ListMoviesFunc: func() ([]models.Movie, error) {
				array := make([]models.Movie, 0)
				return array, nil
			},
		}}

	http.HandlerFunc(controller.ListMovies).ServeHTTP(rec, req)

	expectedStatusCode := http.StatusNoContent
	if expectedStatusCode != rec.Code {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedStatusCode, rec.Code)
	}
}

func TestMovieController_ListMovies_ShouldReturnError_WhenError(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/movies", nil)

	var controller = &controllers.MovieController{
		DB: &mock.MovieDaoMock{
			ListMoviesFunc: func() ([]models.Movie, error) {
				return nil, errors.New("Error listing Movies")
			},
		}}

	http.HandlerFunc(controller.ListMovies).ServeHTTP(rec, req)

	expectedBody := "Error listing Movies"
	expectedStatusCode := http.StatusInternalServerError
	if expectedBody != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedBody, rec.Body.String())
	}
	if expectedStatusCode != rec.Code {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedStatusCode, rec.Code)
	}
}

func TestMovieController_GetMovie_ShouldReturnMovie(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/movies/abc", nil)

	var controller = &controllers.MovieController{
		DB: &mock.MovieDaoMock{
			GetMovieByIdFunc: func(id string) (models.Movie, error) {
				movie := &models.Movie{Id: "abc", Title: "Title", Year: 1999}
				return *movie, nil
			},
		}}

	http.HandlerFunc(controller.GetMovie).ServeHTTP(rec, req)

	var resultMovie models.Movie

	json.Unmarshal(rec.Body.Bytes(), &resultMovie)

	movieTitle := resultMovie.Title

	expectedStatusCode := http.StatusOK
	expectedTitle := "Title"

	if expectedTitle != movieTitle {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedTitle, movieTitle)
	}
	if expectedStatusCode != rec.Code {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedStatusCode, rec.Code)
	}
}

func TestMovieController_GetMovie_ShouldReturnNotFound_WhenNotFound(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/movies/abc", nil)

	var controller = &controllers.MovieController{
		DB: &mock.MovieDaoMock{
			GetMovieByIdFunc: func(id string) (models.Movie, error) {
				return models.Movie{}, nil
			},
		}}

	http.HandlerFunc(controller.GetMovie).ServeHTTP(rec, req)

	expectedStatusCode := http.StatusNotFound

	if expectedStatusCode != rec.Code {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedStatusCode, rec.Code)
	}
}

func TestMovieController_GetMovie_ShouldReturnError_WhenError(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/movies/abc", nil)

	var controller = &controllers.MovieController{
		DB: &mock.MovieDaoMock{
			GetMovieByIdFunc: func(id string) (models.Movie, error) {
				return models.Movie{}, errors.New("Error")
			},
		}}

	http.HandlerFunc(controller.GetMovie).ServeHTTP(rec, req)

	expectedStatusCode := http.StatusInternalServerError
	expectedBody := "Error getting Movie"

	if expectedBody != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedBody, rec.Body.String())
	}
	if expectedStatusCode != rec.Code {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedStatusCode, rec.Code)
	}
}

func TestMovieController_UpdateMovie_ShouldReturnNotFound_WhenNotFound(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("UPDATE", "/movies/abc", nil)

	var controller = &controllers.MovieController{
		DB: &mock.MovieDaoMock{
			GetMovieByIdFunc: func(id string) (models.Movie, error) {
				return models.Movie{}, nil
			},
		}}

	http.HandlerFunc(controller.UpdateMovie).ServeHTTP(rec, req)

	expectedStatusCode := http.StatusNotFound

	if expectedStatusCode != rec.Code {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedStatusCode, rec.Code)
	}
}

func TestMovieController_UpdateMovie_ShouldReturnError_WhenError(t *testing.T) {
	movie := &models.Movie{Title: "Newtitle", Year: 2000}
	movieJson, _ := json.Marshal(movie)
	reqBody := bytes.NewBuffer(movieJson)
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("UPDATE", "/movies/abc", reqBody)

	var controller = &controllers.MovieController{
		DB: &mock.MovieDaoMock{
			UpdateMovieFunc: func(movie, newMovie models.Movie) (models.Movie, error) {
				return models.Movie{}, errors.New("Error")
			},
			GetMovieByIdFunc: func(id string) (models.Movie, error) {
				return models.Movie{Id: "abc", Title: "OldTitle", Year: 1980}, nil
			},
		}}

	http.HandlerFunc(controller.UpdateMovie).ServeHTTP(rec, req)

	expectedStatusCode := http.StatusInternalServerError
	expectedBody := "Error updating Movie"

	if expectedBody != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedBody, rec.Body.String())
	}
	if expectedStatusCode != rec.Code {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedStatusCode, rec.Code)
	}
}

func TestMovieController_UpdateMovie_ShouldReturnUpdatedMovie(t *testing.T) {
	movie := &models.Movie{Title: "Newtitle", Year: 2000}
	movieJson, _ := json.Marshal(movie)
	reqBody := bytes.NewBuffer(movieJson)
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("UPDATE", "/movies/abc", reqBody)

	var controller = &controllers.MovieController{
		DB: &mock.MovieDaoMock{
			UpdateMovieFunc: func(movie, newMovie models.Movie) (models.Movie, error) {
				return models.Movie{Id: "abc", Title: "NewTitle", Year: 2000}, nil
			},
			GetMovieByIdFunc: func(id string) (models.Movie, error) {
				return models.Movie{Id: "abc", Title: "OldTitle", Year: 1980}, nil
			},
		}}

	http.HandlerFunc(controller.UpdateMovie).ServeHTTP(rec, req)

	var resultMovie models.Movie

	json.Unmarshal(rec.Body.Bytes(), &resultMovie)

	movieTitle := resultMovie.Title

	expectedStatusCode := http.StatusOK
	expectedTitle := "NewTitle"

	if expectedTitle != movieTitle {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedTitle, movieTitle)
	}
	if expectedStatusCode != rec.Code {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedStatusCode, rec.Code)
	}
}

func TestMovieController_DeleteMovie_ShouldReturnError_WhenError(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/movies/abc", nil)

	var controller = &controllers.MovieController{
		DB: &mock.MovieDaoMock{
			DeleteMovieFunc: func(id string) error {
				return errors.New("Error")
			},
		}}

	http.HandlerFunc(controller.DeleteMovie).ServeHTTP(rec, req)

	expectedStatusCode := http.StatusInternalServerError
	expectedBody := "Error deleting Movie"

	if expectedBody != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedBody, rec.Body.String())
	}

	if expectedStatusCode != rec.Code {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedStatusCode, rec.Code)
	}
}

func TestMovieController_DeleteMovie_ShouldReturnOK_WhenOK(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/movies/abc", nil)

	var controller = &controllers.MovieController{
		DB: &mock.MovieDaoMock{
			DeleteMovieFunc: func(id string) error {
				return nil
			},
		}}

	http.HandlerFunc(controller.DeleteMovie).ServeHTTP(rec, req)

	expectedStatusCode := http.StatusOK

	if expectedStatusCode != rec.Code {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedStatusCode, rec.Code)
	}
}
