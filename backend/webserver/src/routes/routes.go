package routes

import (
	"go-sandbox/src/controllers"
	"go-sandbox/src/database/factory"
	"log"
	"os"

	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router *mux.Router) {
	dbEngine := os.Getenv("DB_ENGINE")
	log.Println("Creating moviedao with", dbEngine, "as engine")

	var movieController = &controllers.MovieController{DB: factory.MovieDao(dbEngine)}

	router.HandleFunc("/hello", controllers.Hello).Methods("GET")
	router.HandleFunc("/hello/{helloId}", controllers.HelloFromInterface).Methods("GET")
	router.HandleFunc("/auth", controllers.Authorize).Methods("POST")
	router.HandleFunc("/movies", movieController.PostMovie).Methods("POST")
	router.HandleFunc("/movies", movieController.ListMovies).Methods("GET")
	router.HandleFunc("/movies/{movieId}", movieController.GetMovie).Methods("GET")
	router.HandleFunc("/movies/{movieId}", movieController.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movies/{movieId}", movieController.DeleteMovie).Methods("DELETE")
}
