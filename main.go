package main

import (
	"github.com/gorilla/mux"
	controller "github.com/igorthebarros/mux-restapi/Controller"
	"log"
	"net/http"
)

func main() {

	bookController := &controller.BookController{}

	r := mux.NewRouter()
	r.HandleFunc("/api/books", bookController.List).Methods("GET")
	r.HandleFunc("/api/books/{id}", bookController.Find).Methods("GET")
	r.HandleFunc("/api/books", bookController.Create).Methods("POST")
	r.HandleFunc("/api/books/{id}", bookController.Update).Methods("PUT")
	r.HandleFunc("/api/books/{id}", bookController.Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}
