package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tutorials/go/crud/pkg/db"
	"github.com/tutorials/go/crud/pkg/handlers"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()
	// router.HandleFunc("/books", h.GetAllBooks).Methods(http.MethodGet)
	// router.HandleFunc("/books/{id}", h.GetBook).Methods(http.MethodGet)
	// router.HandleFunc("/books", h.AddBook).Methods(http.MethodPost)
	// router.HandleFunc("/books/{id}", h.UpdateBook).Methods(http.MethodPut)
	// router.HandleFunc("/books/{id}", h.DeleteBook).Methods(http.MethodDelete)
	router.HandleFunc("/Product", h.GetProducts).Methods(http.MethodGet)
	log.Println("API is running!")
	http.ListenAndServe("127.0.0.1:3000", router)
}
