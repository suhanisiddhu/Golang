package main

import (
	"github.com/gorilla/mux"
	"layeredProject/datastore/book"
	"layeredProject/http/bookhttp"
	"layeredProject/service/bookservice"
	"log"

	"layeredProject/datastore/author"
	"layeredProject/driver"
	"layeredProject/http/authorhttp"
	"layeredProject/service/authorservice"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	Db := driver.Connection()
	defer Db.Close()

	authorStore := author.New(Db)
	authorService := authorservice.New(authorStore)
	authorHandler := authorhttp.New(authorService)

	r.HandleFunc("/author", authorHandler.PostAuthor).Methods("POST")
	r.HandleFunc("/author/{id}", authorHandler.PutAuthor).Methods("PUT")
	r.HandleFunc("/author/{id}", authorHandler.DeleteAuthor).Methods("DELETE")

	bookStore := book.New(Db)
	bookService := bookservice.New(bookStore)
	bookHandler := bookhttp.New(bookService)

	r.HandleFunc("/book", bookHandler.GetAllBook).Methods("GET")
	r.HandleFunc("/book/{id}", bookHandler.GetBookByID).Methods("GET")
	r.HandleFunc("/book", bookHandler.PostBook).Methods("POST")
	r.HandleFunc("/book/{id}", bookHandler.PutBook).Methods("PUT")
	r.HandleFunc("/book/{id}", bookHandler.DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
