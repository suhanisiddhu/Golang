package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"layeredProject/datastore/author"
	"layeredProject/datastore/book"
	"layeredProject/driver"
	"layeredProject/http/authorhttp"
	"layeredProject/http/bookhttp"
	"layeredProject/service/authorservice"
	"layeredProject/service/bookservice"
)

func main() {
	r := mux.NewRouter()

	DB := driver.Connection()
	defer DB.Close()

	authorStore := author.New(DB)
	authorService := authorservice.New(authorStore)
	authorHandler := authorhttp.New(authorService)

	r.HandleFunc("/author", authorHandler.PostAuthor).Methods("POST")
	r.HandleFunc("/author/{id}", authorHandler.PutAuthor).Methods("PUT")
	r.HandleFunc("/author/{id}", authorHandler.DeleteAuthor).Methods("DELETE")

	bookStore := book.New(DB)
	bookService := bookservice.New(bookStore)
	bookHandler := bookhttp.New(bookService)

	r.HandleFunc("/book", bookHandler.GetAllBook).Methods("GET")
	r.HandleFunc("/book/{id}", bookHandler.GetBookByID).Methods("GET")
	r.HandleFunc("/book", bookHandler.PostBook).Methods("POST")
	r.HandleFunc("/book/{id}", bookHandler.PutBook).Methods("PUT")
	r.HandleFunc("/book/{id}", bookHandler.DeleteBook).Methods("DELETE")

	server := http.Server{
		Addr:    ":8000",
		Handler: r,
	}

	fmt.Println("server started")

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println(err)
	}
}

//log.Fatal(http.ListenAndServe(":8000", r))
