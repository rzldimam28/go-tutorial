package handler

import (
	"encoding/json"
	"math/rand"
	"movie-crud/entity"
	"movie-crud/helper"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// mock data
var movies = []entity.Movie{
	{ID: "1", Title: "Spiderman", Actor: &entity.Actor{FirstName: "Tom", LastName: "Holland"}},
	{ID: "2", Title: "Doctor Strange", Actor: &entity.Actor{FirstName: "Benedict", LastName: "Cumberbatch"}},
}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(&movies)
	helper.LogError(err)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, movie := range movies {
		if movie.ID == params["movieId"] {
			err := json.NewEncoder(w).Encode(&movie)
			helper.LogError(err)
		}
	}
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie entity.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	helper.LogError(err)
	movie.ID = strconv.Itoa(rand.Intn(1000))
	movies = append(movies, movie)
	err = json.NewEncoder(w).Encode(&movie)
	helper.LogError(err)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, movie := range movies {
		if movie.ID == params["movieId"] {
			movies = append(movies[:i], movies[i+1:]...)
			var movie entity.Movie
			err := json.NewDecoder(r.Body).Decode(&movie)
			helper.LogError(err)
			movie.ID = params["movieId"]
			movies = append(movies, movie)
			err = json.NewEncoder(w).Encode(&movie)
			helper.LogError(err)
		}
	}
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, movie := range movies {
		if movie.ID == params["movieId"] {
			movies = append(movies[:i], movies[i+1:]...)
			err := json.NewEncoder(w).Encode(&movies)
			helper.LogError(err)
		}
	}
}