package main

import (
	"fmt"
	"movie-crud/handler"
	"movie-crud/helper"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// mux router
	router := mux.NewRouter()

	// route
	router.HandleFunc("/movies", handler.GetMovies).Methods("GET")
	router.HandleFunc("/movies/{movieId}", handler.GetMovie).Methods("GET")
	router.HandleFunc("/movies", handler.CreateMovie).Methods("POST")
	router.HandleFunc("/movies/{movieId}", handler.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movies/{movieId}", handler.DeleteMovie).Methods("DELETE")

	fmt.Println("Server is running on port 3000...")

	// run server
	err := http.ListenAndServe("127.0.0.1:3000", router)
	helper.LogError(err)

}