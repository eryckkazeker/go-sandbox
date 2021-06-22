package models

type Movie struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Year  int    `json:"year"`
}
