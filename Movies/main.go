package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isdn     string    `json:"isdn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicatio/json")
	params := mux.Vars(r)
	for index, movie := range movies {
		if movie.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, movie := range movies {
		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}

	}

}
func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{
		ID:    "1",
		Isdn:  "4344",
		Title: "Movie one",
		Director: &Director{
			FirstName: "Sanjay",
			LastName:  "bhansali",
		},
	})

	movies = append(movies, Movie{
		ID:    "2",
		Isdn:  "434",
		Title: "Movie two",
		Director: &Director{
			FirstName: "Karan",
			LastName:  "Johar",
		},
	})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", getMovieById).Methods("GET")
	r.HandleFunc("/movie", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	fmt.Println("Starting server at port :8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
