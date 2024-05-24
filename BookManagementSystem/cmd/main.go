package main

import (
	"log"
	"net/http"

	"github.com/Kratos-28/BookManagementSystem/pkg/routes"
	_ "github.com/Kratos-28/BookManagementSystem/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
