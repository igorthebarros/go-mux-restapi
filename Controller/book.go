package Controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/igorthebarros/mux-restapi/Model"
	"math/rand"
	"net/http"
	"strconv"
)

type BookController struct {
}

var books []Model.Book

func (book BookController) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		return
	}
}

func (book BookController) Find(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for _, item := range books {
		if item.Id == params["id"] {
			err := json.NewEncoder(w).Encode(item)
			if err != nil {
				return
			}
			return
		}
	}

	err := json.NewEncoder(w).Encode(&Model.Book{})
	if err != nil {
		return
	}
}

func (book BookController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var bookModel Model.Book

	_ = json.NewDecoder(r.Body).Decode(&bookModel)
	bookModel.Id = strconv.Itoa(rand.Intn(100))
	books = append(books, bookModel)

	err := json.NewEncoder(w).Encode(book)
	if err != nil {
		return
	}
}

func (book BookController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.Id == params["id"] {
			books = append(books[:index], books[index+1:]...)

			var book Model.Book

			_ = json.NewDecoder(r.Body).Decode(&book)
			book.Id = strconv.Itoa(rand.Intn(100))
			books = append(books, book)

			err := json.NewEncoder(w).Encode(book)
			if err != nil {
				return
			}
		}
	}
	json.NewEncoder(w).Encode(books)
}

func (book BookController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.Id == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}
