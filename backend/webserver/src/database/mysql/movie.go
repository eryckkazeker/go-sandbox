package mysql

import (
	"database/sql"
	"go-sandbox/src/models"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type MySqlMovieDaoImpl struct {
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func (dao MySqlMovieDaoImpl) AddMovie(movie models.Movie) (string, error) {
	db := dbConn()
	id := uuid.New().String()
	insert, err := db.Prepare("INSERT INTO movies(id, title, year) values(?,?,?)")
	if err != nil {
		return "", err
	}
	_, insertError := insert.Exec(id, movie.Title, movie.Year)
	if insertError != nil {
		return "", insertError
	}
	//FIXME Find out how to return the id of the inserted movie
	defer db.Close()
	return id, nil
}

func (dao MySqlMovieDaoImpl) ListMovies() ([]models.Movie, error) {
	returnList := make([]models.Movie, 0)
	db := dbConn()
	sel, err := db.Query("SELECT * FROM movies")
	if err != nil {
		return nil, err
	}

	for sel.Next() {
		var id, title string
		var year int
		err = sel.Scan(&id, &title, &year)
		if err != nil {
			return nil, err
		}
		movie := models.Movie{Id: id, Title: title, Year: year}
		returnList = append(returnList, movie)
	}
	defer db.Close()
	return returnList, nil
}

func (dao MySqlMovieDaoImpl) GetMovieById(id string) (models.Movie, error) {
	db := dbConn()
	defer db.Close()
	returnMovie := models.Movie{}
	sel, err := db.Query("SELECT * FROM movies WHERE id = ?", id)
	if err != nil {
		return models.Movie{}, err
	}

	for sel.Next() {
		var id, title string
		var year int
		err = sel.Scan(&id, &title, &year)
		if err != nil {
			return models.Movie{}, err
		}
		returnMovie.Id = id
		returnMovie.Title = title
		returnMovie.Year = year
	}

	return returnMovie, nil

}

func (dao MySqlMovieDaoImpl) UpdateMovie(movie, newMovie models.Movie) (models.Movie, error) {
	db := dbConn()
	update, err := db.Prepare("UPDATE movies SET title=?, year=? WHERE id=?")
	if err != nil {
		return models.Movie{}, err
	}
	_, updateError := update.Exec(newMovie.Title, newMovie.Year, movie.Id)
	if updateError != nil {
		return models.Movie{}, updateError
	}
	defer db.Close()
	return newMovie, nil
}

func (dao MySqlMovieDaoImpl) DeleteMovie(id string) error {
	db := dbConn()
	del, err := db.Prepare("DELETE FROM movies WHERE id=?")
	if err != nil {
		return err
	}
	del.Exec(id)
	defer db.Close()
	return nil
}
