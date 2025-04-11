package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, movie := range movies {
		if movie.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, movie := range movies {
		if movie.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	http.Error(w, "Movie not found", http.StatusNotFound)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, movie := range movies {
		if movie.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
			var updatedMovie Movie
			_ = json.NewDecoder(r.Body).Decode(&updatedMovie)
			updatedMovie.ID = id
			movies = append(movies, updatedMovie)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedMovie)
			return
		}
	}
	http.Error(w, "Movie not found", http.StatusNotFound)
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "438-1234567890", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "438-0987654321", Title: "Movie Two", Director: &Director{Firstname: "Jane", Lastname: "Smith"}})
	movies = append(movies, Movie{ID: "3", Isbn: "438-1122334455", Title: "Movie Three", Director: &Director{Firstname: "Jim", Lastname: "Brown"}})
	movies = append(movies, Movie{ID: "4", Isbn: "438-5566778899", Title: "Movie Four", Director: &Director{Firstname: "Jake", Lastname: "White"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return
	}
	fmt.Println("Server started successfully")
}
