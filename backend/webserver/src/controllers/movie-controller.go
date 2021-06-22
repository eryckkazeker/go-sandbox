package controllers

import (
	"encoding/json"
	"fmt"
	"go-sandbox/src/database"
	"go-sandbox/src/models"
	"go-sandbox/src/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type MovieController struct {
	DB database.MovieDao
}

func (m MovieController) PostMovie(w http.ResponseWriter, r *http.Request) {
	var movie = &models.Movie{}
	utils.ParseBody(r, movie)
	result, err := m.DB.AddMovie(*movie)
	if err != nil {
		log.Println("Error adding movie", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error saving Movie")
		return
	}
	fmt.Fprint(w, result)
}

func (m MovieController) ListMovies(w http.ResponseWriter, r *http.Request) {
	list, err := m.DB.ListMovies()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error listing Movies")
		return
	}

	if len(list) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	listJson, _ := json.Marshal(list)

	fmt.Fprint(w, string(listJson))
}

func (m MovieController) GetMovie(w http.ResponseWriter, r *http.Request) {
	movie, err := m.getMovieByIdFromRequest(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error getting Movie")
		return
	}
	if movie.Title == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	movieJson, err := json.Marshal(movie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(movieJson))
}

func (m MovieController) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	movie, _ := m.getMovieByIdFromRequest(r)
	log.Println("Movie found:", movie)
	if movie.Title == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var newMovie = &models.Movie{}
	utils.ParseBody(r, newMovie)
	log.Println("New movie:", newMovie)
	response, err := m.DB.UpdateMovie(movie, *newMovie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error updating Movie")
		return
	}
	responseJson, _ := json.Marshal(response)
	fmt.Fprint(w, string(responseJson))
}

func (m MovieController) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	urlVars := mux.Vars(r)
	movieId := urlVars["movieId"]
	err := m.DB.DeleteMovie(movieId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error deleting Movie")
	}
}

func (m MovieController) getMovieByIdFromRequest(r *http.Request) (models.Movie, error) {
	urlVars := mux.Vars(r)
	movieId := urlVars["movieId"]
	movie := new(models.Movie)
	log.Println("looking for id", movieId)
	var err error
	*movie, err = m.DB.GetMovieById(movieId)
	if err != nil {
		return models.Movie{}, err
	}
	return *movie, nil
}
